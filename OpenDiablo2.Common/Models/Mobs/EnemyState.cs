using OpenDiablo2.Common.Enums.Mobs;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class EnemyState : MobState
    {
        public int ExperienceGiven { get; protected set; }

        public EnemyState() : base() { }

        public EnemyState(string name, int id, int level, int maxhealth, float x, float y, int experiencegiven)
            : base(name, id, level, maxhealth, x, y)
        {
            ExperienceGiven = experiencegiven;
            AddFlag(eMobFlags.ENEMY);
        }
    }
}
