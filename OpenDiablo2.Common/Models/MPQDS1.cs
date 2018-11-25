using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Common.Models
{

    public sealed class MPQDS1TileProps
    {
        public byte Prop1 { get; internal set; }
        public byte Prop2 { get; internal set; }
        public byte Prop3 { get; internal set; }
        public byte Prop4 { get; internal set; }
    }

    public sealed class MPQDS1WallOrientationTileProps
    {
        public byte Orientation1 { get; internal set; }
        public byte Orientation2 { get; internal set; }
        public byte Orientation3 { get; internal set; }
        public byte Orientation4 { get; internal set; }
    }

    public sealed class MPQDS1WallLayer
    {
        public MPQDS1TileProps[] Props { get; internal set; }
        public MPQDS1WallOrientationTileProps[] Orientations { get; internal set; }
    }

    public sealed class MPQDS1FloorLayer
    {
        public MPQDS1TileProps[] Props { get; internal set; }
    }

    public sealed class MPQDS1ShadowLayer
    {
        public MPQDS1TileProps[] Props { get; internal set; }
    }

    public sealed class MPQDS1Object
    {
        public Int32 Type { get; internal set; }
        public Int32 Id { get; internal set; }
        public Int32 X { get; internal set; }
        public Int32 Y { get; internal set; }
        public Int32 DS1Flags { get; internal set; }
    }

    public sealed class MPQDS1Group
    {
        public Int32 TileX { get; internal set; }
        public Int32 TileY { get; internal set; }
        public Int32 Width { get; internal set; }
        public Int32 Height { get; internal set; }
    }

    public sealed class MPQDS1
    {
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        public Int32 Version { get; internal set; }
        public Int32 Width { get; internal set; }
        public Int32 Height { get; internal set; }
        public Int32 Act { get; internal set; } // ???
        public Int32 TagType { get; internal set; } // ???
        public Int32 FileCount { get; internal set; }
        public Int32 NumberOfWalls { get; internal set; }
        public Int32 NumberOfFloors { get; internal set; }
        public Int32 NumberOfObjects { get; internal set; }
        public Int32 NumberOfNPCs { get; internal set; }
        public Int32 NumberOfTags { get; internal set; }
        public Int32 NumberOfGroups { get; internal set; }

        public MPQDT1[] DT1s = new MPQDT1[33];

        public List<string> FileNames { get; internal set; } = new List<string>();
        public MPQDS1WallLayer[] WallLayers { get; internal set; }
        public MPQDS1FloorLayer[] FloorLayers { get; internal set; }
        public MPQDS1ShadowLayer[] ShadowLayers { get; internal set; }
        public MPQDS1Object[] Objects { get; internal set; }
        public MPQDS1Group[] Groups { get; internal set; }

        // TODO: DI magic please
        public MPQDS1(Stream stream, LevelPreset level, LevelDetail levelDetail, LevelType levelType, IEngineDataManager engineDataManager, IResourceManager resourceManager)
        {
            log.Debug($"Loading {level.Name} (Act {levelDetail.Act})...");
            var br = new BinaryReader(stream);
            Version = br.ReadInt32();
            Width = br.ReadInt32() + 1;
            Height = br.ReadInt32() + 1;
            Act = br.ReadInt32() + 1;

            if (Version >= 10)
            {
                TagType = br.ReadInt32();
                if (TagType == 1 || TagType == 2)
                    NumberOfTags = 1;
            }

            FileCount = 0;
            if (Version >= 3)
            {
                FileCount = br.ReadInt32();
                for (int i = 0; i < FileCount; i++)
                {
                    var fn = "";
                    while (true)
                    {
                        var b = br.ReadByte();
                        if (b == 0)
                            break;
                        fn += (char)b;
                    }
                    if (fn.StartsWith("\\d2\\"))
                        fn = fn.Substring(4);
                    FileNames.Add(fn);
                }
            }

            if (Version >= 9 && Version <= 13)
            {
                br.ReadBytes(8);
            }

            if (Version >= 4)
            {
                NumberOfWalls = br.ReadInt32();

                if (Version >= 16)
                {
                    NumberOfFloors = br.ReadInt32();
                }
                else NumberOfFloors = 1;
            }
            else
            {
                NumberOfFloors = 1;
                NumberOfWalls = 1;
                NumberOfTags = 1;
            }



            var layoutStream = new List<int>();

            if (Version < 4)
            {
                layoutStream.AddRange(new int[] { 1, 9, 5, 12, 11 });
            }
            else
            {
                for (var x = 0; x < NumberOfWalls; x++)
                {
                    layoutStream.Add(1 + x);
                    layoutStream.Add(5 + x);
                }
                for (var x = 0; x < NumberOfFloors; x++)
                    layoutStream.Add(9 + x);

                layoutStream.Add(11);

                if (NumberOfTags > 0)
                    layoutStream.Add(12);
            }

            WallLayers = new MPQDS1WallLayer[NumberOfWalls];
            for (var l = 0; l < NumberOfWalls; l++)
            {
                WallLayers[l] = new MPQDS1WallLayer
                {
                    Orientations = new MPQDS1WallOrientationTileProps[Width * Height],
                    Props = new MPQDS1TileProps[Width * Height]
                };
            }

            FloorLayers = new MPQDS1FloorLayer[NumberOfFloors];
            for (var l = 0; l < NumberOfFloors; l++)
            {
                FloorLayers[l] = new MPQDS1FloorLayer
                {
                    Props = new MPQDS1TileProps[Width * Height]
                };
            }

            ShadowLayers = new MPQDS1ShadowLayer[1];
            for (var l = 0; l < 1; l++)
            {
                ShadowLayers[l] = new MPQDS1ShadowLayer
                {
                    Props = new MPQDS1TileProps[Width * Height]
                };
            }


            for (int n = 0; n < layoutStream.Count(); n++)
            {
                for (var y = 0; y < Height; y++)
                {
                    for (var x = 0; x < Width; x++)
                    {
                        switch (layoutStream[n])
                        {
                            // Walls
                            case 1:
                            case 2:
                            case 3:
                            case 4:
                                WallLayers[layoutStream[n] - 1].Props[x + (y * Width)] = new MPQDS1TileProps
                                {
                                    Prop1 = br.ReadByte(),
                                    Prop2 = br.ReadByte(),
                                    Prop3 = br.ReadByte(),
                                    Prop4 = br.ReadByte()
                                };
                                break;

                            // Orientations
                            case 5:
                            case 6:
                            case 7:
                            case 8:
                                // TODO: Orientations
                                if (Version < 7)
                                {
                                    br.ReadBytes(4);
                                }
                                else
                                {
                                    br.ReadBytes(4);
                                }
                                break;

                            // Floors
                            case 9:
                            case 10:
                                FloorLayers[layoutStream[n] - 9].Props[x + (y * Width)] = new MPQDS1TileProps
                                {
                                    Prop1 = br.ReadByte(),
                                    Prop2 = br.ReadByte(),
                                    Prop3 = br.ReadByte(),
                                    Prop4 = br.ReadByte()
                                };
                                break;
                            // Shadow
                            case 11:
                                ShadowLayers[layoutStream[n] - 11].Props[x + (y * Width)] = new MPQDS1TileProps
                                {
                                    Prop1 = br.ReadByte(),
                                    Prop2 = br.ReadByte(),
                                    Prop3 = br.ReadByte(),
                                    Prop4 = br.ReadByte()
                                };
                                break;
                            // Tags
                            case 12:
                                // TODO: Tags
                                br.ReadBytes(4);
                                break;
                        }
                    }
                }
            }
            /*
            NumberOfObjects = 0;
            if (Version >= 2)
            {

                NumberOfObjects = br.ReadInt32();
                Objects = new MPQDS1Object[NumberOfObjects];

                for (int i = 0; i < NumberOfObjects; i++)
                {
                    Objects[i] = new MPQDS1Object
                    {
                        Type = br.ReadInt32(),
                        Id = br.ReadInt32(),
                        X = br.ReadInt32(),
                        Y = br.ReadInt32()
                    };

                    if (Version > 5)
                        Objects[i].DS1Flags = br.ReadInt32();

                }
            }


            if (Version >= 12 && TagType == 1 || TagType == 2)
            {
                if (Version >= 18)
                    br.ReadBytes(4);

                NumberOfGroups = br.ReadInt32();
            }
            else NumberOfGroups = 0;


            Groups = new MPQDS1Group[NumberOfGroups];
            for (var x = 0; x < NumberOfGroups; x++)
            {
                Groups[x] = new MPQDS1Group
                {
                    TileX = br.ReadInt32(),
                    TileY = br.ReadInt32(),
                    Width = br.ReadInt32(),
                    Height = br.ReadInt32(),
                };
                if (Version >= 13)
                    br.ReadBytes(4);
            }

            if (Version >= 14)
            {

                NumberOfNPCs = br.ReadInt32();
            }*/




            var dt1Mask = level.Dt1Mask;
            for (int i = 0; i < 32; i++)
            {
                var tilePath = levelType.File[i];
                var isMasked = ((dt1Mask >> i) & 1) == 1;
                if (!isMasked)
                    continue;

                log.Debug($"Loading DT resource {levelType.File[i]}");

                DT1s[i] = resourceManager.GetMPQDT1("data\\global\\tiles\\" + levelType.File[i].Replace("/", "\\"));

            }
        }
    }
}
