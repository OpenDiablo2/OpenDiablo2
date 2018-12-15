using OpenDiablo2.Common.Interfaces.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class MissileTypeConfig : IMissileTypeConfig
    {
        public string Name { get; private set; }
        public int Id { get; private set; }

        public int BaseVelocity { get; private set; }
        public int MaxVelocity { get; private set; }
        public int Acceleration { get; private set; }
        public int Range { get; private set; }
        public int ExtraRangePerLevel { get; private set; }

        public int LightRadius { get; private set; }
        public bool LightFlicker { get; private set; }
        public byte[] LightColor { get; private set; } // rgb

        public int FramesBeforeVisible { get; private set; } // how many anim frames before it becomes visible
        public int FramesBeforeActive { get; private set; } // anim frames before it can do anything, e.g. collide
        public bool LoopAnimation { get; private set; } // true: repeat animation until missile hits range or is otherwise destroyed
        // false: repeat animation once, then vanish (WARNING: only becomes invisible, not deleted)
        public string CelFilePath { get; private set; }
        public int AnimationRate { get; private set; } // does not seem to be used
        public int AnimationLength { get; private set; } // length of animation for one direction
        public int AnimationSpeed { get; private set; } // actually used, frames per second
        public int StartingFrame { get; private set; } // 'randstart' is a misnomer, actually just starts at this frame
        public bool AnimationHasSubLoop { get; private set; } // if true, will repeat in a certain range of frames
        //instead of repeating the whole animation again
        public int AnimationSubLoopStart { get; private set; } // what frame to start at if it has a subloop
        public int AnimationSubLoopEnd { get; private set; } // what frame to end at in subloop
        // when it hits this, goes back to subloop start or goes to end if missile is set to die (out of range)

        public MissileTypeConfig(string Name, int Id,
            int BaseVelocity, int MaxVelocity, int Acceleration, int Range, int ExtraRangePerLevel,
            int LightRadius, bool LightFlicker, byte[] LightColor,
            int FramesBeforeVisible, int FramesBeforeActive, bool LoopAnimation, string CelFilePath,
            int AnimationRate, int AnimationLength, int AnimationSpeed, int StartingFrame, 
            bool AnimationHasSubLoop, int AnimationSubLoopStart, int AnimationSubLoopEnd)
        {
            this.Name = Name;
            this.Id = Id;

            this.BaseVelocity = BaseVelocity;
            this.MaxVelocity = MaxVelocity;
            this.Acceleration = Acceleration;
            this.Range = Range;
            this.ExtraRangePerLevel = ExtraRangePerLevel;

            this.LightRadius = LightRadius;
            this.LightFlicker = LightFlicker;
            this.LightColor = LightColor;

            this.FramesBeforeVisible = FramesBeforeVisible;
            this.FramesBeforeActive = FramesBeforeActive;
            this.LoopAnimation = LoopAnimation;
            this.CelFilePath = CelFilePath;
            this.AnimationRate = AnimationRate;
            this.AnimationLength = AnimationLength;
            this.AnimationSpeed = AnimationSpeed;
            this.StartingFrame = StartingFrame;
            this.AnimationHasSubLoop = AnimationHasSubLoop;
            this.AnimationSubLoopStart = AnimationSubLoopStart;
            this.AnimationSubLoopEnd = AnimationSubLoopEnd;
        }
    }

    public static class MissileTypeConfigHelper
    {
        //Missile Id  Vel MaxVel  Accel Range   LevRange 
        //0	      1	  2	  3	      4	    5	    6
        //Light   Flicker Red Green Blue
        //7	      8	      9	  10	11	
        //InitSteps Activate
        //12	    13	
        //LoopAnim CelFile AnimLen RandStart   SubLoop SubStart    SubStop
        //14	   15	   16	   17	       18	   19	       20	
        //CollideType CollideKill CollideFriend   LastCollide Collision
        //21	      22	      23	          24	      25
        //ClientSend NextHit NextDelay Size    CanDestroy ToHit   AlwaysExplode Explosion
        //26	     27	     28	       29	   30	      31	  32	        33
        //NeverDup ReturnFire  GetHit KnockBack   Trans Qty Pierce
        //34	   35	       36	  37	      38	39	40	   
        //Param1  Param1 Comment  Param2 Param2 Comment SpecialSetup
        //41	  42	          43	 44	            45	    	
        //Open  Beta    Skill HitShift    SrcDamage MinDamage   MaxDamage LevDamage
        //46	47      48    49	      50	    51	        52	      53	    	   
        //EType EMin    Emax ELevel  ELen ELevelLen   HitClass NumDirections   LocalBlood
        //54	55      56	 57	     58	  59	      60	   61	           62


        public static IMissileTypeConfig ToMissileTypeConfig(this string[] row)
        {
            return new MissileTypeConfig(
                Name: row[0],
                Id: Convert.ToInt32(row[1]),

                BaseVelocity: Convert.ToInt32(row[2]),
                MaxVelocity: Convert.ToInt32(row[3]),
                Acceleration: Convert.ToInt32(row[4]),
                Range: Convert.ToInt32(row[5]),
                ExtraRangePerLevel: Convert.ToInt32(row[6]),

                LightRadius: Convert.ToInt32(row[7]),
                LightFlicker: (row[8] == "1"),
                LightColor: new byte[] {Convert.ToByte(row[9]), Convert.ToByte(row[10]), Convert.ToByte(row[11])},

                FramesBeforeVisible: Convert.ToInt32(row[12]),
                FramesBeforeActive: Convert.ToInt32(row[13]),
                LoopAnimation: (row[14] == "1"),
                CelFilePath: row[15],

                // TODO: these rows are wrong! research why our missiles.txt has different columns thatn the one in the guide???
                // TODO: UNFINISHED
                AnimationRate: Convert.ToInt32(row[16]),
                AnimationLength: Convert.ToInt32(row[17]),
                AnimationSpeed: Convert.ToInt32(row[18]),
                StartingFrame: Convert.ToInt32(row[19]),
                AnimationHasSubLoop: (row[20] == "1"),
                AnimationSubLoopStart: Convert.ToInt32(row[16]),
                AnimationSubLoopEnd: Convert.ToInt32(row[16])
                );
        }
    }
}
