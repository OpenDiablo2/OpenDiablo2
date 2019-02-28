#ifndef OPENDIABLO2_COMMON_D2DATAMANAGER_H
#define OPENDIABLO2_COMMON_D2DATAMANAGER_H

#include <memory>
#include <map>
#include <string>
#include <OpenDiablo2.Common/D2EngineConfig.h>

namespace OpenDiablo2::Common {

class D2DataManager {
public:
	typedef std::unique_ptr<D2DataManager> Ptr;
	D2DataManager(const D2EngineConfig& engineConfig);
private:
	std::map<std::string, std::string> fileEntries;
};

}

#endif // OPENDIABLO2_COMMON_D2DATAMANAGER_H
