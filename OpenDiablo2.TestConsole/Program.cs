using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Models;
using OpenDiablo2.Core;

namespace OpenDiablo2.TestConsole
{
    static class Program
    {
        private static GlobalConfiguration GlobalConfig = null;
        private static MpqProvider MPQProv = null;
        private static EngineDataManager EngineDataMan = null;

        static void Main(string[] args)
        {
            Console.WriteLine("OpenDiablo2 TestConsole Loaded");
            Console.WriteLine("Type 'help' for more info, 'exit' to close.");
            bool exit = false;
            while (!exit)
            {
                Console.Write("> ");
                string command = Console.ReadLine();
                List<string> words = GetWords(command);//new List<string>(command.Split(' '));
                string func = words[0].ToUpper();
                words.RemoveAt(0);

                switch (func)
                {
                    case "HELP":
                        Help(words);
                        break;
                    case "EXIT":
                        exit = true;
                        break;
                    case "LOADDATA":
                        LoadData(words);
                        break;
                    case "WRITELEVELIDS":
                        WriteELevelIds(words);
                        break;
                }
            }
            // close out the program after exiting
            Environment.Exit(0);
        }


        public static List<string> GetWords(string msg)
        {
            List<string> outs = new List<string>();

            string curword = "";
            bool inquote = false;
            int i = 0;
            while (i < msg.Length)
            {
                char c = msg[i];
                if (c == '"')
                {
                    inquote = !inquote;
                }
                else if (c == ' ' && !inquote)
                {
                    outs.Add(curword);
                    curword = "";
                }
                else
                {
                    curword += c;
                }
                i++;
            }
            outs.Add(curword);

            return outs;
        }

        public static void Help(List<string> words)
        {
            if(words.Count < 1)
            {
                Console.WriteLine("Commands: HELP, EXIT, LOADDATA, WRITELEVELIDS");
                Console.WriteLine("Type 'HELP [command]' for more info.");
                return;
            }

            string command = words[0].ToUpper();
            switch (command)
            {
                case "LOADDATA":
                    Console.WriteLine("'LOADDATA [path to mpq folder]' will load the MPQs and parse their data.");
                    Console.WriteLine("This function is pre-requisite to many others.");
                    return;
                case "WRITELEVELIDS":
                    Console.WriteLine("'WRITELEVELIDS [filepath to write to]' will write the level ids enum to a file.");
                    Console.WriteLine("This function requires the data to be loaded first.");
                    return;
            }

            // if we get to the end, no command was matched
            Help(new List<string>()); // run with an empty list to trigger default message
        }

        public static void LoadData(List<string> words)
        {
            if (words.Count < 1)
            {
                Console.WriteLine("Must supply a path to the folder where the MPQs are stored.");
                return;
            }
            string path = words[0];

            if (EngineDataMan != null)
            {
                Console.WriteLine("Data already loaded!");
                return;
            }

            try
            {
                GlobalConfig = new GlobalConfiguration
                {
                    BaseDataPath = Path.GetFullPath(path)
                };

                MPQProv = new MpqProvider(GlobalConfig);

                EngineDataMan = new EngineDataManager(MPQProv);
            }
            catch(Exception e)
            {
                Console.WriteLine("Failed to load data! " + e.Message);
                return;
            }

            Console.WriteLine("Data loaded.");
        }

        public static void WriteELevelIds(List<string> words)
        {
            if (words.Count < 1)
            {
                Console.WriteLine("Must supply a filepath to write to.");
                return;
            }
            if(EngineDataMan == null)
            {
                Console.WriteLine("You must load the MPQ data first! See 'HELP LOADDATA'.");
                return;
            }
            string path = words[0];
            string output = "public enum eLevelId\r\n{\r\nNone,\r\n";
            output += ELevelIdHelper.GenerateEnum(EngineDataMan.Levels.Select(x => x.LevelPreset).Distinct().ToList());
            output += "}";
            File.WriteAllText(path, output);
            Console.WriteLine("Wrote eLevelIds enum to " + path + ".");
        }
    }
}
