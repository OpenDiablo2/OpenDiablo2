package d2records

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	testify "github.com/stretchr/testify/assert"
)

// Verify the lookup returns the right object after indexing.
func TestIndexObjects(t *testing.T) {
	assert := testify.New(t)

	r, err := NewRecordManager(d2util.LogLevelDefault)
	if err != nil {
		t.Error(err)
	}

	testObjects := []ObjectLookupRecord{
		{Act: 1, Type: d2enum.ObjectTypeCharacter, ID: 0, Description: "Act1CharID0"},
		{Act: 1, Type: d2enum.ObjectTypeCharacter, ID: 1, Description: "Act1CharID1"},
		{Act: 1, Type: d2enum.ObjectTypeCharacter, ID: 2, Description: "Act1CharID2"},
		{Act: 1, Type: d2enum.ObjectTypeCharacter, ID: 3, Description: "Act1CharID3"},
		{Act: 1, Type: d2enum.ObjectTypeItem, ID: 0, Description: "Act1ItemID0"},
		{Act: 1, Type: d2enum.ObjectTypeItem, ID: 1, Description: "Act1ItemID1"},
		{Act: 1, Type: d2enum.ObjectTypeItem, ID: 2, Description: "Act1ItemID2"},
		{Act: 1, Type: d2enum.ObjectTypeItem, ID: 3, Description: "Act1ItemID3"},
		{Act: 2, Type: d2enum.ObjectTypeCharacter, ID: 0, Description: "Act2CharID0"},
		{Act: 2, Type: d2enum.ObjectTypeCharacter, ID: 1, Description: "Act2CharID1"},
		{Act: 2, Type: d2enum.ObjectTypeCharacter, ID: 2, Description: "Act2CharID2"},
		{Act: 2, Type: d2enum.ObjectTypeCharacter, ID: 3, Description: "Act2CharID3"},
		{Act: 2, Type: d2enum.ObjectTypeItem, ID: 0, Description: "Act2ItemID0"},
		{Act: 2, Type: d2enum.ObjectTypeItem, ID: 1, Description: "Act2ItemID1"},
		{Act: 2, Type: d2enum.ObjectTypeItem, ID: 2, Description: "Act2ItemID2"},
		{Act: 2, Type: d2enum.ObjectTypeItem, ID: 3, Description: "Act2ItemID3"},
	}

	r.initObjectRecords(testObjects)

	typeCharacter := int(d2enum.ObjectTypeCharacter)
	typeItem := int(d2enum.ObjectTypeItem)

	assert.Equal("Act1CharId2", r.lookupObject(1, typeCharacter, 2).Description)
	assert.Equal("Act1ItemId0", r.lookupObject(1, typeItem, 0).Description)
	assert.Equal("Act2CharId3", r.lookupObject(2, typeCharacter, 3).Description)
	assert.Equal("Act2ItemId1", r.lookupObject(2, typeItem, 1).Description)
}
