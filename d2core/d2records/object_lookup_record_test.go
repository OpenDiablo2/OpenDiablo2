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
		{Act: 1, Type: d2enum.ObjectTypeCharacter, Id: 0, Description: "Act1CharId0"},
		{Act: 1, Type: d2enum.ObjectTypeCharacter, Id: 1, Description: "Act1CharId1"},
		{Act: 1, Type: d2enum.ObjectTypeCharacter, Id: 2, Description: "Act1CharId2"},
		{Act: 1, Type: d2enum.ObjectTypeCharacter, Id: 3, Description: "Act1CharId3"},
		{Act: 1, Type: d2enum.ObjectTypeItem, Id: 0, Description: "Act1ItemId0"},
		{Act: 1, Type: d2enum.ObjectTypeItem, Id: 1, Description: "Act1ItemId1"},
		{Act: 1, Type: d2enum.ObjectTypeItem, Id: 2, Description: "Act1ItemId2"},
		{Act: 1, Type: d2enum.ObjectTypeItem, Id: 3, Description: "Act1ItemId3"},
		{Act: 2, Type: d2enum.ObjectTypeCharacter, Id: 0, Description: "Act2CharId0"},
		{Act: 2, Type: d2enum.ObjectTypeCharacter, Id: 1, Description: "Act2CharId1"},
		{Act: 2, Type: d2enum.ObjectTypeCharacter, Id: 2, Description: "Act2CharId2"},
		{Act: 2, Type: d2enum.ObjectTypeCharacter, Id: 3, Description: "Act2CharId3"},
		{Act: 2, Type: d2enum.ObjectTypeItem, Id: 0, Description: "Act2ItemId0"},
		{Act: 2, Type: d2enum.ObjectTypeItem, Id: 1, Description: "Act2ItemId1"},
		{Act: 2, Type: d2enum.ObjectTypeItem, Id: 2, Description: "Act2ItemId2"},
		{Act: 2, Type: d2enum.ObjectTypeItem, Id: 3, Description: "Act2ItemId3"},
	}

	r.initObjectRecords(testObjects)

	typeCharacter := int(d2enum.ObjectTypeCharacter)
	typeItem := int(d2enum.ObjectTypeItem)

	assert.Equal("Act1CharId2", r.lookupObject(1, typeCharacter, 2).Description)
	assert.Equal("Act1ItemId0", r.lookupObject(1, typeItem, 0).Description)
	assert.Equal("Act2CharId3", r.lookupObject(2, typeCharacter, 3).Description)
	assert.Equal("Act2ItemId1", r.lookupObject(2, typeItem, 1).Description)
}
