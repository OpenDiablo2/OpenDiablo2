using System;
using System.Collections.Generic;
using System.Drawing;
using OpenDiablo2.Common.Enums.Mobs;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class MobState : IComparable
    {
        public readonly string Name;
        public readonly int Id;
        public bool Alive { get; protected set; } = true;

        /// <summary>The X tile location of the mob</summary>
        public float X { get; set; } = 0;

        /// <summary>The Y tile location of the mob</summary>
        public float Y { get; set; } = 0;

        /// <summary>The speed of the mob (in units per second)</summary>
        public float MovementSpeed { get; set; }

        /// <summary>Represents the movement direction of the mob (16 angular segments)</summary>
        public int MovementDirection { get; set; }

        protected Stat Health;

        protected Dictionary<eDamageTypes, StatDouble> Resistances = new Dictionary<eDamageTypes, StatDouble>();
        protected List<eDamageTypes> Immunities = new List<eDamageTypes>();
        
        public int Level { get; protected set; }

        protected Dictionary<eMobFlags, bool> Flags = new Dictionary<eMobFlags, bool>();

        public MobState() { }

        public MobState(string name, int id, int level, int maxhealth, float x, float y)
        {
            Name = name;
            Id = id;
            Level = level;
            Health = new Stat(0, maxhealth, maxhealth, true);
            X = x;
            Y = y;
        }

        #region Position and Movement
        public void Move(PointF pos)
        {
            Move(pos.X, pos.Y);
        }
        public void Move(float x, float y)
        {
            X = x;
            Y = y;
        }
        public PointF GetPosition()
        {
            return new PointF(X, Y);
        }
        public float GetDistance(float x, float y)
        {
            // note: does not consider pathfinding!
            return (float)Math.Sqrt(((X - x) * (X - x)) + ((Y - y) * (Y - y)));
        }
        #endregion Position and Movement

        #region Combat and Damage
        public void SetResistance(eDamageTypes damagetype, double val)
        {
            if (!Resistances.ContainsKey(damagetype))
            {
                Resistances.Add(damagetype, new StatDouble(0, 100.0, val, false));
            }
            else
            {
                Resistances[damagetype].SetCurrent(val);
            }
        }
        public void AddImmunitiy(eDamageTypes damagetype)
        {
            if (!Immunities.Contains(damagetype))
            {
                Immunities.Add(damagetype);
            }
        }
        public void RemoveImmunity(eDamageTypes damagetype)
        {
            if (Immunities.Contains(damagetype))
            {
                Immunities.Remove(damagetype);
            }
        }

        public int GetHealth()
        {
            return Health.GetCurrent();
        }
        public int GetHealthMax()
        {
            return Health.GetMax();
        }
        public void RecoverHealth(int health)
        {
            Health.AddCurrent(health);
        }
        public int TakeDamage(int damage, eDamageTypes damagetype, MobState source = null)
        {
            // returns the actual amount of damage taken
            damage = HandleResistances(damage, damagetype, source);
            Health.AddCurrent(-1 * damage);
            int newhp = Health.GetCurrent();
            if(newhp <= 0)
            {
                Die(source);
            }
            return damage;
        }
        protected int HandleResistances(int damage, eDamageTypes damagetype, MobState source = null)
        {
            if(damagetype == eDamageTypes.NONE)
            {
                return damage;
            }
            if (Immunities.Contains(damagetype))
            {
                return 0;
            }
            if (!Resistances.ContainsKey(damagetype))
            {
                return damage;
            }

            // TODO: need to verify 1) is damage integer? and 2) if so, how is this rounding down?
            // e.g. is it always 'round down' / 'round up' or does it use 'math.round'
            damage = (int)(damage * Resistances[damagetype].GetCurrent());
            return damage;
        }

        public void Die(MobState source = null)
        {
            // TODO: how do we want to tackle this?
            Alive = false;
        }
        #endregion Combat and Damage

        #region Flags
        public void AddFlag(eMobFlags flag, bool on = true)
        {
            if (Flags.ContainsKey(flag))
            {
                Flags[flag] = on;
            }
            else
            {
                Flags.Add(flag, on);
            }
        }
        public void RemoveFlag(eMobFlags flag)
        {
            if (Flags.ContainsKey(flag))
            {
                Flags.Remove(flag);
            }
        }
        public bool HasFlag(eMobFlags flag)
        {
            if (Flags.ContainsKey(flag))
            {
                return Flags[flag];
            }
            return false;
        }

        #endregion Flags


        public int CompareTo(object obj) => Id - ((MobState)obj).Id;
        public override bool Equals(object obj) => Id == (obj as MobState)?.Id;
        public static bool operator ==(MobState obj1, MobState obj2) => obj1?.Id == obj2?.Id;
        public static bool operator !=(MobState obj1, MobState obj2) => obj1?.Id != obj2?.Id;
        public static bool operator <(MobState obj1, MobState obj2) => obj1?.Id < obj2?.Id;
        public static bool operator >(MobState obj1, MobState obj2) => obj1?.Id > obj2?.Id;
        public static bool operator <=(MobState obj1, MobState obj2) => obj1?.Id <= obj2?.Id;
        public static bool operator >=(MobState obj1, MobState obj2) => obj1?.Id >= obj2?.Id;
        public override int GetHashCode() => Id.GetHashCode();
    }
}
