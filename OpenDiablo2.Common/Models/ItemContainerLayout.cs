using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common.Enums;

namespace OpenDiablo2.Common.Models
{
    public class ItemContainerLayout
    {
        public string ResourceName { get; internal set; }
        public string PaletteName { get; internal set; } = Palettes.Units;
        public int BaseFrame { get; internal set; } = 0;

        public static Dictionary<eItemContainerType, ItemContainerLayout> Values = new Dictionary<eItemContainerType, ItemContainerLayout>
        {
            {eItemContainerType.Helm,  new ItemContainerLayout { ResourceName = ResourcePaths.HelmGlovePlaceholder,  BaseFrame = 1 } },
            {eItemContainerType.Amulet, new ItemContainerLayout{ ResourceName = ResourcePaths.RingAmuletPlaceholder  } },
            {eItemContainerType.Armor, new ItemContainerLayout { ResourceName = ResourcePaths.ArmorPlaceholder } },
            {eItemContainerType.Weapon, new ItemContainerLayout { ResourceName = ResourcePaths.WeaponsPlaceholder } },
            {eItemContainerType.Belt, new ItemContainerLayout { ResourceName = ResourcePaths.BeltPlaceholder } },
            {eItemContainerType.Ring, new ItemContainerLayout { ResourceName = ResourcePaths.RingAmuletPlaceholder, BaseFrame = 1 } },
            {eItemContainerType.Glove, new ItemContainerLayout { ResourceName = ResourcePaths.HelmGlovePlaceholder } },
            {eItemContainerType.Boots, new ItemContainerLayout { ResourceName = ResourcePaths.BootsPlaceholder } },
        };
    }

}
