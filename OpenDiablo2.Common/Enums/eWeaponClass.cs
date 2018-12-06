using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Enums
{
    public enum eWeaponClass
    {
        None,
        HandToHand,
        Bow,
        OneHandSwing,
        OneHandThrust,
        Staff,
        TwoHandSwing,
        TwoHandThrust,
        Crossbow,
        LeftJabRightSwing,
        LeftJabRightThrust,
        LeftSwingRightSwing,
        LeftSwingRightThrust,
        OneHandToHand,
        TwoHandToHand
    }

    public static class eWeaponClassExtensions
    {
        private static readonly Dictionary<eWeaponClass, string> codes = new Dictionary<eWeaponClass, string>
        {
            {eWeaponClass.None                    , "" },
            {eWeaponClass.HandToHand              , "hth" },
            {eWeaponClass.Bow                     , "bow" },
            {eWeaponClass.OneHandSwing            , "1hs" },
            {eWeaponClass.OneHandThrust           , "1ht" },
            {eWeaponClass.Staff                   , "stf" },
            {eWeaponClass.TwoHandSwing            , "2hs" },
            {eWeaponClass.TwoHandThrust           , "2ht" },
            {eWeaponClass.Crossbow                , "xbw" },
            {eWeaponClass.LeftJabRightSwing       , "1js" },
            {eWeaponClass.LeftJabRightThrust      , "1jt" },
            {eWeaponClass.LeftSwingRightSwing     , "1ss" },
            {eWeaponClass.LeftSwingRightThrust    , "1st" },
            {eWeaponClass.OneHandToHand           , "ht1" },
            {eWeaponClass.TwoHandToHand           , "ht2" }
        };

        public static string ToToken(this eWeaponClass source) => codes[source];
        public static eWeaponClass ToWeaponClass(this string source) => codes.First(x => x.Value.ToUpper() == source.ToUpper()).Key;

    }
}
