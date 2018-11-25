using Microsoft.VisualStudio.TestTools.UnitTesting;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Models;
using System.IO;

namespace OpenDiablo2.Core.UT
{
    [TestClass]
    public class UT_EngineDataManager
    {
        private static readonly string DataPath = @"C:\PutYourMPQsHere\";

        private EngineDataManager LoadData()
        {
            GlobalConfiguration globalconfig = new GlobalConfiguration
            {
                BaseDataPath = Path.GetFullPath(DataPath)
            };

            MPQProvider mpqprov = new MPQProvider(globalconfig);

            EngineDataManager edm = new EngineDataManager(mpqprov);

            return edm;
        }

        [TestMethod]
        public void DataLoadTest()
        {
            EngineDataManager edm = LoadData();

            Assert.IsTrue(edm.LevelDetails.Count > 0);
            Assert.IsTrue(edm.LevelPresets.Count > 0);
            Assert.IsTrue(edm.LevelTypes.Count > 0);
        }

        [TestMethod]
        public void GenerateELevelIdTest()
        {
            EngineDataManager edm = LoadData();

            string output = ELevelIdHelper.GenerateEnum(edm.LevelPresets);

            Assert.IsTrue(output.Length > 0);
        }
    }
}
