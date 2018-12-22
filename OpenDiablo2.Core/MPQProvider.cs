/*  OpenDiablo 2 - An open source re-implementation of Diablo 2 in C#
 *  
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>. 
 */

using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Exceptions;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using OpenDiablo2.Common.Models.Mobs;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Linq;

namespace OpenDiablo2.Core
{
    public sealed class MpqProvider : IMPQProvider
    {
        private static readonly log4net.ILog _log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);
        
        private readonly IList<MPQ> _mpqs;
        private readonly Dictionary<string, int> _mpqLookup = new Dictionary<string, int>();

        public MpqProvider(GlobalConfiguration globalConfiguration)
        {
            // TODO: Make this less dumb. We need to an external file to configure mpq load order.
            
            _mpqs = Directory
                .EnumerateFiles(globalConfiguration.BaseDataPath, "*.mpq")
                .Where(x => !(Path.GetFileName(x)?.StartsWith("patch", System.StringComparison.InvariantCultureIgnoreCase) ?? false))
                .Select(file => new MPQ(file))
                .ToList();


            if (!_mpqs.Any())
            {
                _log.Fatal("No data files were found! Are you specifying the correct data path?");
                throw new OpenDiablo2Exception("No data files were found.");
            }

            // Load the base game files
            for(var i = 0; i < _mpqs.Count; i++)
            {
                var path = Path.GetFileName(_mpqs[i].Path) ?? string.Empty;
                
                if (path.StartsWith("d2exp", System.StringComparison.InvariantCultureIgnoreCase) || path.StartsWith("d2x", System.StringComparison.InvariantCultureIgnoreCase))
                    continue;

                foreach(var file in _mpqs[i].Files)
                    _mpqLookup[file.ToLower()] = i;
            }

            // Load the expansion game files
            for (var i = 0; i < _mpqs.Count; i++)
            {
                var path = Path.GetFileName(_mpqs[i].Path) ?? string.Empty;
                
                if (!path.StartsWith("d2exp", System.StringComparison.InvariantCultureIgnoreCase) && !path.StartsWith("d2x", System.StringComparison.InvariantCultureIgnoreCase))
                    continue;

                foreach (var file in _mpqs[i].Files)
                    _mpqLookup[file.ToLower()] = i;
            }

            // Get the combined list file by joining all of the other mpqs
            List<string> superListFile = _mpqs.SelectMany(x => x.Files).ToList();

            var patchMPQ = Directory
                .EnumerateFiles(globalConfiguration.BaseDataPath, "*.mpq")
                .Where(x => Path.GetFileName(x).StartsWith("patch", System.StringComparison.InvariantCultureIgnoreCase))
                .Select(file => new MPQ(file, superListFile))
                .First();

            _mpqs.Add(patchMPQ);
            int patchMPQIndex = _mpqs.Count - 1;
            
            // Replace existing mpqLookups with those from the patch, which take precedence
            foreach (var file in patchMPQ.Files)
            {
                // unlike the other mpqs, we need to ensure that the files actually exist
                // inside of the patch mpq instead of assuming that they do, because
                // we can't trust the filelist
                if (!patchMPQ.HasFile(file))
                {
                    continue;
                }

                _mpqLookup[file.ToLower()] = patchMPQIndex;
            }
        }

        public byte[] GetBytes(string fileName)
        {
            var stream = GetStream(fileName);
            var result = new byte[stream.Length];
            stream.Read(result, 0, (int)stream.Length);
            return result;
        }

        public IEnumerator<MPQ> GetEnumerator() => _mpqs.GetEnumerator();

        IEnumerator IEnumerable.GetEnumerator() => GetEnumerator();

        public Stream GetStream(string fileName) => !_mpqLookup.ContainsKey(fileName.ToLower()) 
            ? null 
            : _mpqs[_mpqLookup[fileName.ToLower()]].OpenFile(fileName);

        public IEnumerable<string> GetTextFile(string fileName)
            => new StreamReader(_mpqs[_mpqLookup[fileName.ToLower()]].OpenFile(fileName)).ReadToEnd().Split('\n');

        public string GetCharacterDccPath(eHero hero, eMobMode mobMode, eCompositType compositType, PlayerEquipment equipment)
        {
            var fileName = $@"{ResourcePaths.PlayerAnimationBase}\{hero.ToToken()}\{compositType.ToToken()}\{hero.ToToken()}{compositType.ToToken()}".ToLower();
            var armorType = eArmorType.Lite;

            // Override default armor type based on equipped torso
            if(equipment.Torso != null && (equipment.Torso.Item as Armor).ArmorTypes.ContainsKey(compositType))
                armorType = (equipment.Torso.Item as Armor).ArmorTypes[compositType];

            switch (compositType)
            {
                case eCompositType.Head:
                    fileName += $"{equipment.Head?.Item.Code ?? eArmorType.Lite.ToToken()}{mobMode.ToToken()}";
                    return _mpqLookup.ContainsKey($"{fileName}{equipment.WeaponClass.ToToken()}.dcc".ToLower())
                        ? $"{fileName}{equipment.WeaponClass.ToToken()}.dcc".ToLower()
                        : $"{fileName}{eWeaponClass.HandToHand.ToToken()}.dcc".ToLower();
                case eCompositType.Torso:
                case eCompositType.Legs:
                case eCompositType.RightArm:
                case eCompositType.LeftArm:
                    fileName += $"{armorType.ToToken()}{mobMode.ToToken()}";
                    return _mpqLookup.ContainsKey($"{fileName}{equipment.WeaponClass.ToToken()}.dcc".ToLower())
                        ? $"{fileName}{equipment.WeaponClass.ToToken()}.dcc".ToLower()
                        : $"{fileName}{eWeaponClass.HandToHand.ToToken()}.dcc".ToLower();
                case eCompositType.RightHand:
                    if (!(equipment.RightArm?.Item is Weapon))
                        return null;
                    fileName += $"{equipment.RightArm.Item.Code}{mobMode.ToToken()}{equipment.WeaponClass.ToToken()}.dcc".ToLower();
                    return fileName;
                case eCompositType.LeftHand:
                    if (!(equipment.LeftArm?.Item is Weapon))
                        return null;
                    fileName += $"{equipment.LeftArm.Item.Code}{mobMode.ToToken()}{equipment.WeaponClass.ToToken()}.dcc".ToLower();
                    return fileName;
                case eCompositType.Shield:
                    if (!(equipment.LeftArm?.Item is Armor))
                        return null;
                    fileName += $"{equipment.LeftArm.Item.Code}{mobMode.ToToken()}";
                    return _mpqLookup.ContainsKey($"{fileName}{equipment.WeaponClass.ToToken()}.dcc".ToLower())
                        ? $"{fileName}{equipment.WeaponClass.ToToken()}.dcc".ToLower()
                        : $"{fileName}{eWeaponClass.HandToHand.ToToken()}.dcc".ToLower();
                // TODO: Figure these out...
                case eCompositType.Special1:
                case eCompositType.Special2:
                    fileName += $"{armorType.ToToken()}{mobMode.ToToken()}{equipment.WeaponClass}.dcc".ToLower();
                    return _mpqLookup.ContainsKey(fileName)
                        ? fileName
                        : null; // TODO: Should we silence this?
                case eCompositType.Special3:
                case eCompositType.Special4:
                case eCompositType.Special5:
                case eCompositType.Special6:
                case eCompositType.Special7:
                case eCompositType.Special8:
                default:
                    return null;
            }
        }
    }
}
