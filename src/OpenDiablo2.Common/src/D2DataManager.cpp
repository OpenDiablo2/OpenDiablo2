#include <experimental/filesystem>
#include <spdlog/spdlog.h>
#include <StormLib.h>
#include <OpenDiablo2.Common/D2DataManager.h>

namespace fs = std::experimental::filesystem;

OpenDiablo2::Common::D2DataManager::D2DataManager(
	const D2EngineConfig &engineConfig)
: fileEntries()
{
	spdlog::info("Loading data files");
	auto mpqExt = std::string(".mpq");
	for (auto &p : fs::recursive_directory_iterator(engineConfig.BasePath))
	{
		if (p.path().extension() != mpqExt)
			continue;

		HANDLE hMpq = NULL;
		HANDLE hListFile = NULL;

		if(!SFileOpenArchive(p.path().c_str(), 0, 0, &hMpq)) {
			spdlog::error(std::string("  > ").append(p.path().string()).append(" [READ ERROR!]"));
			continue;
		}

		if(!SFileOpenFileEx(hMpq, "(listfile)", 0, &hListFile)) {
			spdlog::error(std::string("  > ").append(p.path().string()).append(" [LIST FILE NOT FOUND!]"));
			SFileCloseArchive(hMpq);
			continue;
		}

		auto listFileContents = std::string();
		char szBuffer[0x10000];
		DWORD dwBytes = 1;

		while(dwBytes > 0)
		{
			SFileReadFile(hListFile, szBuffer, sizeof(szBuffer), &dwBytes, NULL);
			if(dwBytes > 0) {
				listFileContents.append(szBuffer);
			}
		}

		std::string delim = "\r\n";
		auto start = 0U;
		auto end = listFileContents.find(delim);
		auto linesFound = 0;
		while (end != std::string::npos)
		{
			linesFound++;
			fileEntries.emplace(listFileContents.substr(start, end - start), p.path().stem().string());
			start = end + delim.length();
			end = listFileContents.find(delim, start);
		}
		spdlog::debug(std::string("  > ").append(p.path().string()).append(" [").append(std::to_string(linesFound)).append(" files]"));
		SFileCloseFile(hListFile);

		SFileCloseArchive(hMpq);
	}
}
