using OpenDiablo2.Common.Enums.Mobs;
using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class MobState
    {
        public readonly string Name;
        public readonly int Id;
        public bool Alive { get; protected set; } = true;

        protected float X = 0;
        protected float Y = 0;

        protected Stat Health;
        protected Stat Stamina;
        protected Stat Mana;

        protected Dictionary<eMobFlags, bool> Flags = new Dictionary<eMobFlags, bool>();

        public MobState(string name, int id, int maxhealth, int maxmana, int maxstamina, float x, float y)
        {
            Name = name;
            Id = id;
            Health = new Stat(0, maxhealth, maxhealth, true);
            Mana = new Stat(0, maxmana, maxmana, true);
            Stamina = new Stat(0, maxstamina, maxstamina, true);
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

        #region Mana
        public int GetMana()
        {
            return Mana.GetCurrent();
        }
        public int GetManaMax()
        {
            return Mana.GetMax();
        }
        public void RecoverMana(int mana)
        {
            Mana.AddCurrent(mana);
        }
        public void UseMana(int mana)
        {
            Mana.AddCurrent(-mana);
        }
        #endregion Mana

        #region Stamina
        public int GetStamina()
        {
            return Stamina.GetCurrent();
        }
        public int GetStaminaMax()
        {
            return Stamina.GetMax();
        }
        public void RecoverStamina(int stamina)
        {
            Stamina.AddCurrent(stamina);
        }
        public void UseStamina(int stamina)
        {
            Stamina.AddCurrent(-stamina);
        }
        #endregion Stamina

        #region Combat and Damage
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
            // TODO: implement resistances based on damage type and change 'damage'
            Health.AddCurrent(-1 * damage);
            int newhp = Health.GetCurrent();
            if(newhp <= 0)
            {
                Die(source);
            }
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
    }
}
