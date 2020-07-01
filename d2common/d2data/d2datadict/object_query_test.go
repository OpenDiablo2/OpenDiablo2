package d2datadict

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	testify "github.com/stretchr/testify/assert"
)

// Verify the lookup returns the right object after indexing.
func TestIndexObjects(t *testing.T) {
	assert := testify.New(t)

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

	indexedTestObjects := indexObjects(testObjects)

	typeCharacter := int(d2enum.ObjectTypeCharacter)
	typeItem := int(d2enum.ObjectTypeItem)

	assert.Equal("Act1CharId2", lookupObject(1, typeCharacter, 2, indexedTestObjects).Description)
	assert.Equal("Act1ItemId0", lookupObject(1, typeItem, 0, indexedTestObjects).Description)
	assert.Equal("Act2CharId3", lookupObject(2, typeCharacter, 3, indexedTestObjects).Description)
	assert.Equal("Act2ItemId1", lookupObject(2, typeItem, 1, indexedTestObjects).Description)
}
