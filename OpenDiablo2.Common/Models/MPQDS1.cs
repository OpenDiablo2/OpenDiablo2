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

        public MPQDT1[] DT1s = new MPQDT1[32];
        
        public List<string> FileNames { get; internal set; } = new List<string>();
        public List<MPQDS1WallLayer> WallLayers { get; internal set; } = new List<MPQDS1WallLayer>();
        public List<MPQDS1FloorLayer> FloorLayers { get; internal set; } = new List<MPQDS1FloorLayer>();
        public MPQDS1ShadowLayer ShadowLayer { get; internal set; } = new MPQDS1ShadowLayer();
        public List<MPQDS1Object> Objects { get; internal set; } = new List<MPQDS1Object>();

        // TODO: DI magic please
        public MPQDS1(Stream stream, string fileName, int definition, int act, IEngineDataManager engineDataManager, IResourceManager resourceManager)
        {
            log.Debug($"Loading {fileName} (Act {act}) Def {definition}");
            var br = new BinaryReader(stream);
            Version = br.ReadInt32();
            Width = br.ReadInt32() + 1;
            Height = br.ReadInt32() + 1;
            Act = br.ReadInt32();
            TagType = br.ReadInt32();
            FileCount = br.ReadInt32();



            if (TagType != 0)
                throw new ApplicationException("We don't currently handle those tag things...");

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

            NumberOfWalls = br.ReadInt32();
            NumberOfFloors = br.ReadInt32();

            for (int i = 0; i < NumberOfWalls; i++)
            {
                var wallLayer = new MPQDS1WallLayer
                {
                    Props = new MPQDS1TileProps[Width * Height],
                    Orientations = new MPQDS1WallOrientationTileProps[Width * Height]
                };

                for (int y = 0; y < Height; y++)
                {
                    for (int x = 0; x < Width; x++)
                    {
                        wallLayer.Props[x + (y * Width)] = new MPQDS1TileProps
                        {
                            Prop1 = br.ReadByte(),
                            Prop2 = br.ReadByte(),
                            Prop3 = br.ReadByte(),
                            Prop4 = br.ReadByte()
                        };
                    }
                }

                for (int y = 0; y < Height; y++)
                {
                    for (int x = 0; x < Width; x++)
                    {
                        wallLayer.Orientations[x + (y * Width)] = new MPQDS1WallOrientationTileProps
                        {
                            Orientation1 = br.ReadByte(),
                            Orientation2 = br.ReadByte(),
                            Orientation3 = br.ReadByte(),
                            Orientation4 = br.ReadByte()
                        };
                    }
                }
                WallLayers.Add(wallLayer);
            }

            for (int i = 0; i < NumberOfFloors; i++)
            {
                var floorLayer = new MPQDS1FloorLayer
                {
                    Props = new MPQDS1TileProps[Width * Height]
                };

                for (int y = 0; y < Height; y++)
                {
                    for (int x = 0; x < Width; x++)
                    {
                        floorLayer.Props[x + (y * Width)] = new MPQDS1TileProps
                        {
                            Prop1 = br.ReadByte(),
                            Prop2 = br.ReadByte(),
                            Prop3 = br.ReadByte(),
                            Prop4 = br.ReadByte()
                        };
                    }
                }
                FloorLayers.Add(floorLayer);
            }

            ShadowLayer.Props = new MPQDS1TileProps[Width * Height];
            for (int y = 0; y < Height; y++)
            {
                for (int x = 0; x < Width; x++)
                {
                    ShadowLayer.Props[x + (y * Width)] = new MPQDS1TileProps
                    {
                        Prop1 = br.ReadByte(),
                        Prop2 = br.ReadByte(),
                        Prop3 = br.ReadByte(),
                        Prop4 = br.ReadByte()
                    };
                }
            }

            // TODO: Tag layer goes here (tag = 1)

            NumberOfObjects = br.ReadInt32();
            for (int i = 0; i < NumberOfObjects; i++)
            {
                Objects.Add(new MPQDS1Object
                {
                    Type = br.ReadInt32(),
                    Id = br.ReadInt32(),
                    X = br.ReadInt32(),
                    Y = br.ReadInt32(),
                    DS1Flags = br.ReadInt32()
                });
            }

            // TODO: Option groups go here (tag = 1)

            NumberOfNPCs = br.ReadInt32();


            // TODO: WalkPaths

            LevelPreset levelPreset;
            if (definition == -1)
            {
                levelPreset = engineDataManager.LevelPresets.First(x =>
                   x.File1.ToLower() == fileName.ToLower()
                || x.File2.ToLower() == fileName.ToLower()
                || x.File3.ToLower() == fileName.ToLower()
                || x.File4.ToLower() == fileName.ToLower()
                || x.File5.ToLower() == fileName.ToLower()
                || x.File6.ToLower() == fileName.ToLower());
            }
            else
            {
                levelPreset = engineDataManager.LevelPresets.First(x => x.Def == definition);
            }

            var dt1Mask = levelPreset.Dt1Mask;
            var levelType = engineDataManager.LevelTypes.First(x => x.Id == levelPreset.LevelId && x.Act == act);

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
