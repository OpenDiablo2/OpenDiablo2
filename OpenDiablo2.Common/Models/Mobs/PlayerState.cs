using OpenDiablo2.Common.Enums;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class PlayerState : MobState
    {
        public eHero HeroType { get; protected set; }

        // Player character stats
        protected Stat Vitality;
        protected Stat Strength;
        protected Stat Magic;
        protected Stat Dexterity;


        public PlayerState(string name, int id, int maxhealth, int maxmana, int maxstamina, float x, float y,
            int vitality, int strength, int magic, int dexterity, eHero herotype)
            : base(name, id, maxhealth, maxmana, maxstamina, x, y)
        {
            Vitality = new Stat(0, vitality, vitality, true);
            Strength = new Stat(0, strength, strength, true);
            Magic = new Stat(0, magic, magic, true);
            Dexterity = new Stat(0, dexterity, dexterity, true);

            HeroType = herotype;
        }

        // TODO: when a player equips an item, apply the relevant modifiers to their stats
        // TODO: when a player unequips an item, remove the relevant modifiers from their stats
    }
}
