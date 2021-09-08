package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __000001_init_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xb6\xe6\x02\x04\x00\x00\xff\xff\x2d\x73\xd0\xe1\x1e\x00\x00\x00")

func _000001_init_down_sql() ([]byte, error) {
	return bindata_read(
		__000001_init_down_sql,
		"000001_init.down.sql",
	)
}

var __000001_init_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xd0\x5d\x6f\xb2\x30\x14\x07\xf0\x6b\xf8\x14\xe7\x86\x00\x09\x7a\xf3\x3c\x31\x8b\xbb\xaa\x5a\x27\x1b\x2f\x06\xea\xd4\xdd\x28\x4a\xd9\x9a\x81\x22\xad\xc9\x4c\xd3\xef\xbe\xe0\xd0\x38\x75\x77\x3d\xa7\xcd\xaf\xe7\xfc\xfb\x11\x46\x04\x03\x41\x3d\x0f\x83\x3b\x84\x20\x24\x80\x67\x6e\x4c\x62\x90\xb2\x5d\x56\x34\x63\x5f\x4a\xad\xf2\xed\xfa\x93\x83\xa5\x6b\x2c\x85\x57\x14\xf5\x47\x28\xb2\xfe\x75\x6c\x47\xd7\xa4\x64\x19\xb4\xcb\x2d\x17\xef\x15\xe5\x4a\xb1\x0d\xa7\x95\x58\x24\x02\x88\xeb\xe3\x98\x20\x7f\x4c\xde\x8e\x6c\x30\xf1\x3c\x18\xe0\x21\x9a\x78\x04\x82\x70\x6a\xd9\x8e\x94\x74\x93\x2a\x75\x52\xf8\x2e\x67\x82\x5e\x1a\x03\x44\x70\xed\xdc\x00\x56\x4c\xa2\x61\x7d\x63\x99\xc6\xbc\x65\x14\x2d\x23\x05\x63\xd4\x35\xfc\xae\x91\x99\x0e\x98\x41\x38\x35\xed\x9b\x0f\x8a\x03\xdf\xe5\xf7\x7c\xab\x63\xdf\x9f\xb1\x73\x61\x94\x49\x45\x37\x62\xf1\x47\x04\x8d\xbd\xe4\xeb\x0f\x5a\x24\x4b\x29\x69\xce\xa9\x52\x3f\x65\x63\x40\xcf\x7d\x72\x03\xe2\xe8\x9a\x38\x94\x14\x08\x9e\x1d\xcf\x4c\xe4\xe7\x22\x63\x34\x4f\x39\x5c\xc7\xfa\x1c\x87\xc1\x89\xac\x5f\x36\xa0\xa3\x6b\xeb\x8a\x26\x82\xd6\xcb\x9c\xf1\x7d\x99\x5e\xb7\x52\x9a\xd3\xab\xd6\x38\x72\x7d\x14\xcd\xe1\x05\xcf\xc1\x62\xa9\x03\xe7\x58\x6c\xdd\x86\x5f\x4b\x9d\x12\xa9\x97\x46\x7d\x82\x23\x88\x31\x81\xbd\xc8\x1e\x8a\xd5\xff\x66\x94\x47\xfd\x3b\x00\x00\xff\xff\x60\xc4\xab\x56\x4c\x02\x00\x00")

func _000001_init_up_sql() ([]byte, error) {
	return bindata_read(
		__000001_init_up_sql,
		"000001_init.up.sql",
	)
}

var __000002_system_settings_table_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x2d\xae\x2c\x2e\x49\xcd\x8d\x2f\x4e\x2d\x29\xc9\xcc\x4b\x2f\xb6\xe6\x02\x04\x00\x00\xff\xff\xd2\x63\x5d\x39\x27\x00\x00\x00")

func _000002_system_settings_table_down_sql() ([]byte, error) {
	return bindata_read(
		__000002_system_settings_table_down_sql,
		"000002_system_settings_table.down.sql",
	)
}

var __000002_system_settings_table_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xcc\xc1\x6a\x83\x30\x18\x07\xf0\xb3\x79\x8a\xff\x51\x41\xc4\xc1\x0e\x83\x9d\x32\xf7\xc9\x64\x6e\x2d\xf1\x6b\xd1\x53\x69\x31\x96\x80\x4a\xdb\xc4\x52\x09\x79\xf7\xd2\x17\xf8\x15\x8a\x24\x13\x58\x7e\xd5\x84\xaa\xc4\xff\x86\x41\x6d\xd5\x70\x03\xef\xb3\xcb\x4d\x0f\xe6\x11\x82\x5d\xad\xd3\xd3\xc1\x6a\xe7\xcc\x7c\xb6\x88\x45\x64\x7a\xec\xa5\x2a\x7e\xa4\x8a\xdf\xf2\x3c\x49\x45\x74\x3f\x8e\x8b\x06\x53\xcb\xa9\x88\xb6\xaa\xfa\x93\xaa\xc3\x2f\x75\x88\x4d\x9f\x88\x04\xde\x9b\x01\xd9\xb4\xda\xeb\x18\xc2\x37\x95\x72\x57\x33\x5e\x80\x2c\x98\x14\x1a\x62\x2c\x6e\xf8\x98\x4e\xef\xde\xeb\xb9\x0f\xe1\x53\x3c\x03\x00\x00\xff\xff\x3e\xa0\x26\x35\x9e\x00\x00\x00")

func _000002_system_settings_table_up_sql() ([]byte, error) {
	return bindata_read(
		__000002_system_settings_table_up_sql,
		"000002_system_settings_table.up.sql",
	)
}

var __000003_blocks_rootid_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xe6\x72\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xca\xcf\x2f\x89\xcf\x4c\xb1\xe6\x02\x04\x00\x00\xff\xff\x51\xe5\xe2\x3a\x33\x00\x00\x00")

func _000003_blocks_rootid_down_sql() ([]byte, error) {
	return bindata_read(
		__000003_blocks_rootid_down_sql,
		"000003_blocks_rootid.down.sql",
	)
}

var __000003_blocks_rootid_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xe6\x72\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xca\xcf\x2f\x89\xcf\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x36\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\xc2\x68\x66\x83\x3e\x00\x00\x00")

func _000003_blocks_rootid_up_sql() ([]byte, error) {
	return bindata_read(
		__000003_blocks_rootid_up_sql,
		"000003_blocks_rootid.up.sql",
	)
}

var __000004_auth_table_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x2d\x2d\x4e\x2d\x2a\xb6\xe6\xc2\x2e\x59\x9c\x5a\x5c\x9c\x99\x9f\x57\x6c\xcd\x05\x08\x00\x00\xff\xff\xb6\xc1\x44\xa1\x3d\x00\x00\x00")

func _000004_auth_table_down_sql() ([]byte, error) {
	return bindata_read(
		__000004_auth_table_down_sql,
		"000004_auth_table.down.sql",
	)
}

var __000004_auth_table_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x91\x4d\x4f\x32\x31\x14\x85\xd7\xcc\xaf\xb8\x4b\x48\x08\xe1\x25\x2f\x89\x89\xab\x82\x45\x47\x11\x4c\xa7\x1a\x58\x4d\x2a\xbd\xa3\x8d\xf3\x65\x6f\xf1\x23\x4d\xff\xbb\x99\x48\x24\x61\x30\x71\xa1\x5d\x3e\xa7\xed\xb9\xf7\x9c\xa9\xe0\x4c\x72\x90\x6c\x32\xe7\x10\xcf\x60\xb1\x94\xc0\x57\x71\x22\x13\xf0\x7e\x50\x5b\xcc\xcc\x5b\x08\x5b\x42\x4b\xd0\x8d\x3a\x46\xc3\x1d\x13\xd3\x0b\x26\xba\xff\x86\xc3\x5e\x3f\xea\x34\x52\xa9\x0a\x3c\xe4\x58\x28\x93\x7f\xc1\xd1\x78\xdc\xc0\x5a\x11\xbd\x56\xb6\xf5\x49\x91\xa9\x94\x70\x63\xd1\x1d\x2a\x6a\xeb\x1e\x53\x42\xfb\x62\x36\x7b\x8b\xd1\x5e\xd2\xca\xa9\x96\x8b\xad\x6a\x82\xcf\xe3\xbd\xc9\x60\x50\x57\xe4\x1e\x2c\x52\x08\x97\xc9\x72\xe1\x3d\xe6\x84\x21\x48\xbe\x92\xde\x63\xa9\x43\xe8\x47\x9d\x8d\x45\xe5\x30\x55\xae\x79\x36\x89\xcf\xe3\x85\x6c\xd6\xab\xf5\x11\xaa\x31\xc7\x36\xbd\x11\xf1\x35\x13\x6b\xb8\xe2\x6b\xe8\x1a\xdd\x8b\x7a\x3b\xfb\xe2\x9d\x9e\xf3\x10\xce\xf8\x8c\xdd\xce\x25\x34\xb3\xb2\xa9\xe4\x02\x12\x2e\x61\xeb\xb2\x93\xe2\xfe\xff\x6e\x90\xd3\x28\xfa\x59\x25\x84\x44\xa6\x2a\xbf\x69\xc5\x55\x4f\x58\x1e\xab\x2a\x6d\xdf\xfd\xfb\xb8\x7e\x27\x98\x8f\x00\x00\x00\xff\xff\x43\xa6\x60\x6a\xab\x02\x00\x00")

func _000004_auth_table_up_sql() ([]byte, error) {
	return bindata_read(
		__000004_auth_table_up_sql,
		"000004_auth_table.up.sql",
	)
}

var __000005_blocks_modifiedby_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xe6\x72\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\xcd\x4f\xc9\x4c\xcb\x4c\x4d\x89\x4f\xaa\xb4\xe6\x02\x04\x00\x00\xff\xff\x6a\xfe\x38\x0a\x37\x00\x00\x00")

func _000005_blocks_modifiedby_down_sql() ([]byte, error) {
	return bindata_read(
		__000005_blocks_modifiedby_down_sql,
		"000005_blocks_modifiedby.down.sql",
	)
}

var __000005_blocks_modifiedby_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xe6\x72\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\xcd\x4f\xc9\x4c\xcb\x4c\x4d\x89\x4f\xaa\x54\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x36\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\x30\x55\xd2\xd8\x42\x00\x00\x00")

func _000005_blocks_modifiedby_up_sql() ([]byte, error) {
	return bindata_read(
		__000005_blocks_modifiedby_up_sql,
		"000005_blocks_modifiedby.up.sql",
	)
}

var __000006_sharing_table_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x2d\xce\x48\x2c\xca\xcc\x4b\xb7\xe6\x02\x04\x00\x00\xff\xff\x7a\x74\xe5\xab\x1f\x00\x00\x00")

func _000006_sharing_table_down_sql() ([]byte, error) {
	return bindata_read(
		__000006_sharing_table_down_sql,
		"000006_sharing_table.down.sql",
	)
}

var __000006_sharing_table_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcc\x41\x4b\xc3\x30\x18\x06\xe0\x73\xf3\x2b\xde\x63\x0b\x65\x4c\x14\x11\x3c\xa5\xf5\x9b\x06\x6b\x2b\xe9\xa7\xb8\xd3\x68\x49\xaa\xc1\xb5\x9b\x5b\x06\x8e\x90\xff\x2e\xbd\x78\xd8\xfd\xe1\x29\x35\x49\x26\xb0\x2c\x2a\x82\x5a\xa1\x6e\x18\xf4\xa1\x5a\x6e\x11\xc2\x62\x7f\xb0\x83\xfb\x8d\xf1\xf8\xd5\x1d\xdc\xf4\x89\x54\x24\xce\xe0\x5d\xea\xf2\x49\xea\xf4\xfa\x36\xcb\x45\x62\xa7\xae\xdf\x5a\x83\xa2\x69\x2a\x92\x75\x2e\x12\xbf\xfb\xb6\xd3\xbf\xba\x5a\x2e\x67\x36\xee\x8c\x1b\x9c\x35\x9b\xfe\x7c\x11\x9c\xf6\xa6\xf3\x76\xd3\x79\x14\xea\x51\xd5\x9c\x8b\xe4\x55\xab\x17\xa9\xd7\x78\xa6\x35\x52\x67\x32\x91\x21\x04\x37\x60\x31\x9e\x8f\x3f\xdb\x18\x1f\x68\x25\xdf\x2a\xc6\xbc\xc8\x92\x49\xa3\x25\xc6\xc9\x0f\x77\x63\x7f\x13\x82\x9d\x4c\x8c\xf7\xe2\x2f\x00\x00\xff\xff\x82\xbb\xda\xde\xdc\x00\x00\x00")

func _000006_sharing_table_up_sql() ([]byte, error) {
	return bindata_read(
		__000006_sharing_table_up_sql,
		"000006_sharing_table.up.sql",
	)
}

var __000007_workspaces_table_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x2d\xcf\x2f\xca\x2e\x2e\x48\x4c\x4e\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\x1a\xe4\xe6\x36\x22\x00\x00\x00")

func _000007_workspaces_table_down_sql() ([]byte, error) {
	return bindata_read(
		__000007_workspaces_table_down_sql,
		"000007_workspaces_table.down.sql",
	)
}

var __000007_workspaces_table_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcd\x41\x4b\xc3\x30\x18\xc6\xf1\x73\xf3\x29\xde\xe3\x0a\x65\x4c\x14\x11\x3c\x65\x35\xd3\x6a\xed\x24\xcd\x64\x3b\x95\xce\xbc\x2d\x61\x6b\x1b\x9b\x14\x1d\xe1\xfd\xee\x32\x19\x1e\x3c\x3f\x0f\xff\x5f\x2a\x05\x57\x02\x14\x5f\xe6\x02\xb2\x15\x14\x6b\x05\x62\x9b\x95\xaa\x84\x10\xe6\x76\xc4\xc6\x7c\x13\x7d\x0d\xe3\xc1\xd9\xfa\x03\x1d\xcc\x58\x64\x34\xbc\x73\x99\x3e\x71\x39\xbb\xbe\x8d\x13\x16\x39\xd3\xf6\x93\xad\xfc\x70\xc0\xfe\x6f\xba\x5a\x2c\xe2\xdf\x5c\xb1\xc9\xf3\xf3\x09\xbd\x37\x7d\xeb\x20\x04\xd3\xc0\xdc\x0e\xce\xb7\x23\x3a\xa2\xe7\x72\x5d\x84\x80\x47\x87\x44\x4a\x6c\x55\x08\xd8\x6b\xa2\x84\x45\xdd\xa0\x4d\x63\x50\x57\xfb\xd3\x3f\x71\xb2\xba\xf6\x58\xd5\x1e\x96\xd9\x63\x56\xa8\x84\x45\x6f\x32\x7b\xe5\x72\x07\x2f\x62\x07\x33\xa3\x63\x16\x5f\xa4\xee\xe4\x3e\x8f\x44\x0f\x62\xc5\x37\xb9\x82\x73\x85\xa7\x4a\x48\x28\x85\x82\xc9\x37\x77\xdd\xfe\xe6\x62\xde\xb3\x9f\x00\x00\x00\xff\xff\x29\x38\x4d\xc3\x10\x01\x00\x00")

func _000007_workspaces_table_up_sql() ([]byte, error) {
	return bindata_read(
		__000007_workspaces_table_up_sql,
		"000007_workspaces_table.up.sql",
	)
}

var __000008_teams_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xe6\x72\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xcf\x2f\xca\x2e\x2e\x48\x4c\x4e\x8d\xcf\x4c\xb1\xe6\xe2\xc2\xa1\xb1\x38\x23\xb1\x28\x33\x2f\x9d\x1c\x9d\xa9\xc5\xc5\x99\xf9\x79\xa8\x96\x26\x96\x96\x64\xc4\x17\xa7\x16\x95\x65\x26\xa7\x5a\x73\x01\x02\x00\x00\xff\xff\x24\x48\xc4\xb6\xad\x00\x00\x00")

func _000008_teams_down_sql() ([]byte, error) {
	return bindata_read(
		__000008_teams_down_sql,
		"000008_teams.down.sql",
	)
}

var __000008_teams_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xe6\x72\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xcf\x2f\xca\x2e\x2e\x48\x4c\x4e\x8d\xcf\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x36\xd3\xb4\xe6\xe2\xc2\x61\x46\x71\x46\x62\x51\x66\x5e\x3a\x85\x86\xa4\x16\x17\x67\xe6\xe7\xa1\x38\x25\xb1\xb4\x24\x23\xbe\x38\xb5\xa8\x2c\x33\x39\x15\x6e\x8a\x91\x01\xc8\x94\xd0\x00\x17\xc7\x10\x2c\x3e\x51\x08\x76\x0d\x41\xb5\xdd\x56\x41\xdd\x40\x5d\x21\xdc\xc3\x35\xc8\x15\x43\x42\x5d\xc1\x3f\x08\x55\xd0\x33\x58\xc1\x2f\xd4\xc7\xc7\x9a\x0b\x10\x00\x00\xff\xff\xab\x8d\x48\xa9\x30\x01\x00\x00")

func _000008_teams_up_sql() ([]byte, error) {
	return bindata_read(
		__000008_teams_up_sql,
		"000008_teams.up.sql",
	)
}

var __000009_blocks_history_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xb6\xe6\x72\xf4\x09\x71\x0d\xc2\x25\x1d\x9f\x91\x59\x5c\x92\x5f\x54\xa9\x10\xe4\xea\xe7\xe8\xeb\xaa\x10\xe2\x8f\xcd\x08\x40\x00\x00\x00\xff\xff\x38\xe5\xec\x7a\x61\x00\x00\x00")

func _000009_blocks_history_down_sql() ([]byte, error) {
	return bindata_read(
		__000009_blocks_history_down_sql,
		"000009_blocks_history.down.sql",
	)
}

var __000009_blocks_history_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x93\xd1\x6f\xda\x30\x10\xc6\x9f\xc9\x5f\x71\x2f\x11\xc9\x94\x56\x93\x36\xa1\xa9\x48\x93\x42\x38\xc0\x5b\x62\x57\x8e\xbb\x96\xbd\x50\x20\xce\xb0\x1a\x48\x1a\xa7\xea\x50\x94\xff\x7d\x4a\x47\x28\x0d\xa0\x3d\x4c\x7d\xc3\x3e\xdf\xef\xbe\xfb\xf8\xe2\xfa\x02\x39\x08\x77\xe0\x23\x94\xe5\x65\x96\xcb\x58\xfd\xae\xaa\x45\x92\x2e\x1f\x34\x70\xa4\x6e\x80\x20\xd8\x71\x6d\xb6\x52\xba\x48\xf3\x6d\xdf\xf0\x38\xba\x02\x77\x0c\x32\x02\xca\x04\xe0\x1d\x09\x45\x78\x82\x68\x19\x1d\x15\xc1\x0f\x97\x7b\x13\x97\x5b\x9f\x7a\xb6\x63\x74\xca\x52\xc5\x70\x99\xa5\xba\xf8\x95\x4b\x5d\x55\x6a\xa3\x65\x5e\xcc\xe6\x05\x08\x12\x60\x28\xdc\xe0\x5a\xfc\x7c\xc1\xd2\x1b\xdf\x87\x21\x8e\xdc\x1b\x5f\x00\x65\xb7\x96\xed\x94\xa5\xdc\x44\x55\xd5\x50\xf4\x63\xa2\x0a\x79\xc8\x18\xba\x02\x6b\xce\x11\xc0\x0a\x05\x1f\xd5\x15\xab\x6b\x4e\x2f\xcc\xf5\x85\x19\x81\x39\xb9\x32\x83\x2b\x33\xee\x3a\xd0\xa5\xec\xb6\x6b\x1f\x0d\x58\x6f\xf5\x63\x72\x8a\x6f\xf5\xec\xd3\x1a\x7b\x07\x8c\x6c\x9e\xcb\x4d\x31\x3b\x63\xc1\x8e\x7d\xaf\x97\x2b\xb9\x9e\xdf\x97\xa5\x4c\xb4\xac\xaa\xbf\xc7\x1d\x03\x06\x64\x4c\xa8\x70\x8c\x4e\xb1\xcd\x24\x08\xbc\x7b\xf9\xad\x8a\x64\x7f\x88\x95\x4c\x22\x0d\x6d\x5b\xbf\x85\x8c\x36\xc8\xfa\xe5\x0e\xe8\x18\x9d\x65\x2e\xe7\x85\xac\x97\xd9\xc3\x9f\xb2\xa8\x7d\x15\xc9\x44\xb6\xae\xf2\x34\x3d\xb1\xcc\x3a\x8d\x54\xac\x64\x34\x5b\x6c\x5b\x95\xe7\x34\x7f\xd0\xd9\x7c\x29\x8f\x9b\xae\x39\x09\x5c\x3e\x85\xef\x38\x05\xeb\xf0\x9d\xa3\x22\xdb\xb0\xe1\x8d\x43\x8d\xbd\x75\xbf\xeb\xd5\x09\x0e\x51\xc0\x53\x11\x7f\x59\x2f\x3e\xef\xf6\xea\x1b\xc6\x9b\x1e\x83\xd0\x10\xb9\x00\x32\xa6\x8c\x23\x10\x7a\x2a\xd5\x60\x85\xe8\xa3\x27\xe0\x03\x8c\x38\x0b\xce\xc7\x1e\x18\x1f\x22\x87\xc1\x14\x0e\x92\x80\xa1\x67\xf7\x8d\xe6\xcf\x6e\xfb\xbf\x17\xf0\x4e\x93\x81\x51\xf0\x18\x1d\xf9\xc4\x13\x30\x64\x75\x18\x27\x84\x8e\xdb\x82\x9a\x2f\xa4\x91\xc3\xf8\x3f\x2c\xf9\x4f\x5d\xaf\xf3\x87\xe8\xa3\xc0\x33\x18\x78\x5e\xc9\x5c\xc2\x6b\xc8\xbe\xc2\xc7\xbe\xf1\x27\x00\x00\xff\xff\xd3\x97\xb4\x77\x9f\x04\x00\x00")

func _000009_blocks_history_up_sql() ([]byte, error) {
	return bindata_read(
		__000009_blocks_history_up_sql,
		"000009_blocks_history.up.sql",
	)
}

var __000010_blocks_created_by_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x54\xc1\x6e\x9b\x40\x10\xbd\xf3\x15\x73\x41\x86\xca\x89\x2a\xb5\xb2\xaa\x58\xaa\x84\x61\x6c\xd3\xc2\xae\xb5\x6c\x9a\xb8\x17\xc7\x36\x4b\x8d\x82\x0d\x01\xa2\xd4\x42\xfc\x7b\x85\x0d\x0e\x8e\xb1\xa3\xaa\xce\xa5\xca\x71\x77\x67\xde\x7b\x33\x3b\xf3\x34\x8b\x23\x03\xae\xf5\x2c\x84\x2c\xbb\x8c\x62\xe1\xf9\xbf\xf3\x7c\x16\x84\xf3\xfb\x04\x18\x12\xcd\x46\xe0\xf4\xf0\x6d\x12\x06\x6e\x57\xd2\x19\x6a\x1c\xcb\x7c\xb3\x0f\x84\x72\xc0\x5b\xd3\xe1\x4e\x03\x9a\x22\x01\x00\xf8\x2e\xfc\xd0\x98\x3e\xd4\x98\xf2\xa9\xa3\xb6\x37\x77\x59\xe6\x7b\x70\x19\x85\x49\xfa\x2b\x16\x49\x9e\xfb\xab\x44\xc4\xe9\x64\x9a\x02\x37\x6d\x74\xb8\x66\x8f\xf8\xcf\x0d\x38\xb9\xb6\x2c\x30\xb0\xaf\x5d\x5b\x1c\x08\xbd\x51\xd4\x76\x96\x89\x95\x9b\xe7\x35\xa0\xe4\x21\xf0\x53\x51\x87\x31\x34\x8e\x05\xd4\x01\x86\xe2\x70\xd6\x2f\x5e\x94\x96\x3c\xbe\x90\x97\x17\xb2\x0b\xf2\xf0\x4a\xb6\xaf\x64\xaf\xd5\x86\x16\xa1\x37\x2d\xb5\x89\x63\xb9\x4e\x1e\x82\x26\x0a\xa5\xa3\x36\x2b\xed\xec\xc3\x44\xd3\x58\xac\xd2\xc9\xf1\x76\x94\x0c\x77\xc9\x7c\x21\x96\xd3\xbb\x2c\x13\x41\x22\xf2\x7c\x7b\x2c\x91\xa0\x67\x0e\x4c\xc2\xb7\x69\xe9\x3a\x12\xc0\xf1\xb6\x3a\xfa\x69\x50\x3f\x7b\xbe\x08\xdc\xe4\xa0\xd7\xdf\x1c\x4a\x2a\xec\x22\xb8\x44\xde\xe6\xcc\x63\x31\x4d\x45\x51\x5e\x9d\xe8\x31\x72\x1b\x6e\x5d\x11\x88\xc3\xdb\x38\x0c\x9b\x8b\x5c\x86\xae\xef\xf9\xc2\x9d\xcc\xd6\x87\x8f\x4f\x61\x7c\x9f\x44\xd3\xb9\x68\x4c\x1d\x31\xd3\xd6\xd8\x18\xbe\xe3\x18\x94\x7a\x68\xdb\x77\xd5\x4d\x84\xba\xdf\xc3\xea\x1b\x0a\x18\x4d\x2f\x26\xde\x41\x0e\x8f\xa9\xf7\x65\x39\xfb\x5c\x16\xdc\x95\xa4\xbd\x1c\xc9\x24\x0e\x32\x0e\xe6\x80\x50\x86\x60\x92\xa6\x2d\x00\xc5\x41\x0b\x75\x0e\x1f\xa0\xcf\xa8\xdd\xbc\x26\x40\x99\x81\x0c\x7a\x63\xa8\x4d\x0b\x3a\xba\xda\x95\xaa\x81\x78\xf9\x29\x3b\xf2\x37\x60\x05\x4a\x40\xa7\xa4\x6f\x99\x3a\x07\x83\x16\xc3\x3a\x34\xc9\xe0\xa5\x98\x6a\x89\x2a\x29\x94\xbd\xd2\x8a\x7f\xd0\xf4\xcc\x6d\xa0\x85\x1c\x8f\x40\xc0\xd3\x42\xc4\xa2\x36\x68\x5f\xe1\x63\x57\x32\x18\x1d\x1d\x73\xb0\xad\x4b\x49\xd2\x49\x9b\x9b\x2c\xfc\x24\x0d\xe3\xf5\x49\xbb\x2b\x63\xfe\xde\xf6\x76\xe8\xef\xf6\xf7\x6e\x7f\xff\x99\xfd\x3d\xcf\xf6\x6b\xcb\x5f\x5b\x9f\xb3\xdb\xe1\x79\x55\xbc\x91\x3d\xee\x44\x9e\x41\x63\xcd\x2e\x4f\x79\xdf\x9e\x65\xfd\x09\x00\x00\xff\xff\x5f\xbd\x27\xe2\xe9\x09\x00\x00")

func _000010_blocks_created_by_down_sql() ([]byte, error) {
	return bindata_read(
		__000010_blocks_created_by_down_sql,
		"000010_blocks_created_by.down.sql",
	)
}

var __000010_blocks_created_by_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8e\x41\x4b\xc4\x30\x10\x85\xef\xfd\x15\x73\x6b\x0a\xb2\x20\x82\x97\x65\x0f\xb3\x49\x44\x21\x6e\x25\x4d\x05\x4f\x65\xb7\x9d\xb2\xc1\xc6\x48\x12\xd0\x52\xfa\xdf\xa5\x78\x11\xd4\xb2\x97\x39\xcc\x7b\xef\xe3\x43\x65\xa4\x06\x83\x7b\x25\x61\x9a\x36\xef\x81\x7a\xfb\x39\xcf\xa7\xc1\xb7\xaf\x11\x50\x08\xe0\xa5\xaa\x1f\x0f\xd0\x06\x3a\x26\xea\x9a\xd3\x08\xcf\xa8\xf9\x3d\x6a\x76\x73\x5b\x6c\xb3\x55\x40\x73\xb6\x31\xf9\x30\x5e\x02\xca\xea\x27\x81\xe6\x2f\x8b\x4a\x9a\x9f\xab\x1d\xf0\x12\x95\xac\xb8\x64\x87\x5a\xa9\x87\x3b\xc6\x22\x0d\xd4\x26\x70\xbe\xb3\xbd\xfd\x6e\xf5\xc1\xbb\x15\xa1\x8f\x33\x05\xfa\x3f\xdf\xd8\x0e\x76\xbf\xe3\xe5\x5d\x6a\x21\x35\xec\x5f\xd6\xc6\x6f\x91\x42\x6a\x8e\x09\xb0\xe2\x30\x58\x67\x13\x5c\x17\x57\x90\xe7\xcb\x89\x63\x4c\xe4\xf2\x62\x9b\x7d\x05\x00\x00\xff\xff\xcb\x4a\xd7\xe3\x7d\x01\x00\x00")

func _000010_blocks_created_by_up_sql() ([]byte, error) {
	return bindata_read(
		__000010_blocks_created_by_up_sql,
		"000010_blocks_created_by.up.sql",
	)
}

var __000011_match_collation_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x31\x0e\x85\x20\x0c\x06\xe0\x9d\x53\xfc\x17\xe8\xf0\xe6\x37\x1a\x36\xe3\x22\x17\x80\x58\xa1\xb1\x69\x8d\x74\xf1\xf6\x7e\x44\xd8\x3c\x86\x58\x47\x38\x1a\xe3\x70\x63\x0c\x7e\x38\x11\xa1\x0c\x99\xb8\x6b\x67\xc8\x84\x58\xb0\x85\xb8\x55\xd5\x17\xca\x67\xa0\x69\xb5\x2b\xed\x79\xcd\x4b\xc1\xef\x9f\xbe\x00\x00\x00\xff\xff\x54\x6c\x45\x67\x4e\x00\x00\x00")

func _000011_match_collation_down_sql() ([]byte, error) {
	return bindata_read(
		__000011_match_collation_down_sql,
		"000011_match_collation.down.sql",
	)
}

var __000011_match_collation_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x95\xcd\x6e\xda\x40\x14\x85\xf7\x3c\xc5\x51\x36\x80\xd4\x58\xea\x3a\x42\xaa\x63\xa6\xea\xc2\x0d\x29\xb8\x4a\x77\x68\x62\xae\xf1\x28\xf3\x43\xe6\x8e\x45\x91\xe5\x77\xaf\x3a\xc6\x01\x55\xd9\x51\x29\x6c\xed\x73\xbe\x3b\xdf\xcc\xe2\xb6\xad\xaa\x20\xed\x06\x89\x39\xf0\xab\x46\xb2\xd3\xcd\x56\xd9\xae\x1b\x01\xc0\xed\x2d\x4a\xa7\xb5\x0c\xca\x59\xb8\x0a\x46\x86\x40\xde\x38\x0e\x63\x46\x56\x4b\x6b\x49\x33\x82\x7c\xd6\x14\xf3\x2b\x51\xe0\xcb\x29\x94\xbd\x75\x67\x98\xac\x44\x2e\xb2\xa2\x0f\xaf\x4f\xd4\xca\x3b\x03\x65\x2b\xe7\x4d\xfc\xb0\xe6\xb2\x26\x23\x93\x98\x63\x3c\x7d\x13\x4b\x71\x2c\x59\x69\x08\x33\x8c\x87\xc1\x63\xa4\x0f\xf3\xe3\xbf\xbe\x75\x36\x67\x9e\x16\xe9\x7d\xba\x12\x93\xe9\x74\x7a\x37\x1a\x6c\x9e\xb5\x2b\x5f\xf8\x74\xd6\x66\xb7\x91\x81\xde\xce\xf9\xa3\x21\x7f\xc0\x0c\xd9\xe2\x21\x4b\x8b\xc9\x38\xcd\x0b\xb1\x44\x91\xde\xe7\x02\x6d\x9b\xec\x3c\x55\xea\x77\xd7\xf5\x14\x64\x8b\x3c\x4f\x0b\x81\xf1\xa7\x77\xa5\xa7\x77\x71\xce\xe3\x52\x3c\xa6\x4b\x01\x0e\x26\xe0\xeb\x72\xf1\xfd\xfd\xa9\x7d\x58\xfc\x12\xd9\xcf\xa2\x0f\xf7\x5f\xe6\x22\xcd\xf3\x45\xf6\x77\xce\x39\xe9\x5f\x25\xd4\x8a\x83\xf3\x87\xff\xa3\xb6\x3e\xd2\xae\x44\x91\x89\x59\x39\x7b\xf1\xbb\x0d\x9c\x6b\xd1\xaa\xa5\x57\x76\x7b\xb1\x55\x8f\xb9\x16\xa9\x03\x07\x32\x60\x0a\x41\xd9\xed\xe5\x4f\x16\x71\xeb\x01\x77\x25\x92\x0d\x93\xbf\x58\x2d\x42\xae\x44\x68\xef\xfc\x0b\xef\x64\x49\x17\x5b\x9d\x48\x1f\xac\xd6\xb6\xa4\x99\x4e\x8b\xec\x89\x60\x89\x36\x90\x78\x8d\x26\x35\x79\x82\x0b\x35\xf9\xbd\x62\x42\xa8\x09\x46\x6d\x7d\xbf\x95\xf6\x4a\x6b\x78\xe2\x46\x87\xa1\xaf\x2c\xa4\x05\x99\x5d\x38\x1c\x09\xfb\x9a\x6c\xec\xa9\x0a\xa5\xb3\x1b\x15\xab\x8a\x51\x49\xcd\x94\x0c\x45\x71\x56\x29\x65\xc3\xc4\x90\xb8\xe9\xaf\x73\x2f\xb9\x47\xde\x80\xbc\x77\x3e\x39\x5e\x7f\x5c\x61\x9f\xa3\x84\xdd\x74\xdd\xe8\x4f\x00\x00\x00\xff\xff\xed\x3f\x00\x0b\xa7\x07\x00\x00")

func _000011_match_collation_up_sql() ([]byte, error) {
	return bindata_read(
		__000011_match_collation_up_sql,
		"000011_match_collation.up.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"000001_init.down.sql": _000001_init_down_sql,
	"000001_init.up.sql": _000001_init_up_sql,
	"000002_system_settings_table.down.sql": _000002_system_settings_table_down_sql,
	"000002_system_settings_table.up.sql": _000002_system_settings_table_up_sql,
	"000003_blocks_rootid.down.sql": _000003_blocks_rootid_down_sql,
	"000003_blocks_rootid.up.sql": _000003_blocks_rootid_up_sql,
	"000004_auth_table.down.sql": _000004_auth_table_down_sql,
	"000004_auth_table.up.sql": _000004_auth_table_up_sql,
	"000005_blocks_modifiedby.down.sql": _000005_blocks_modifiedby_down_sql,
	"000005_blocks_modifiedby.up.sql": _000005_blocks_modifiedby_up_sql,
	"000006_sharing_table.down.sql": _000006_sharing_table_down_sql,
	"000006_sharing_table.up.sql": _000006_sharing_table_up_sql,
	"000007_workspaces_table.down.sql": _000007_workspaces_table_down_sql,
	"000007_workspaces_table.up.sql": _000007_workspaces_table_up_sql,
	"000008_teams.down.sql": _000008_teams_down_sql,
	"000008_teams.up.sql": _000008_teams_up_sql,
	"000009_blocks_history.down.sql": _000009_blocks_history_down_sql,
	"000009_blocks_history.up.sql": _000009_blocks_history_up_sql,
	"000010_blocks_created_by.down.sql": _000010_blocks_created_by_down_sql,
	"000010_blocks_created_by.up.sql": _000010_blocks_created_by_up_sql,
	"000011_match_collation.down.sql": _000011_match_collation_down_sql,
	"000011_match_collation.up.sql": _000011_match_collation_up_sql,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"000001_init.down.sql": &_bintree_t{_000001_init_down_sql, map[string]*_bintree_t{
	}},
	"000001_init.up.sql": &_bintree_t{_000001_init_up_sql, map[string]*_bintree_t{
	}},
	"000002_system_settings_table.down.sql": &_bintree_t{_000002_system_settings_table_down_sql, map[string]*_bintree_t{
	}},
	"000002_system_settings_table.up.sql": &_bintree_t{_000002_system_settings_table_up_sql, map[string]*_bintree_t{
	}},
	"000003_blocks_rootid.down.sql": &_bintree_t{_000003_blocks_rootid_down_sql, map[string]*_bintree_t{
	}},
	"000003_blocks_rootid.up.sql": &_bintree_t{_000003_blocks_rootid_up_sql, map[string]*_bintree_t{
	}},
	"000004_auth_table.down.sql": &_bintree_t{_000004_auth_table_down_sql, map[string]*_bintree_t{
	}},
	"000004_auth_table.up.sql": &_bintree_t{_000004_auth_table_up_sql, map[string]*_bintree_t{
	}},
	"000005_blocks_modifiedby.down.sql": &_bintree_t{_000005_blocks_modifiedby_down_sql, map[string]*_bintree_t{
	}},
	"000005_blocks_modifiedby.up.sql": &_bintree_t{_000005_blocks_modifiedby_up_sql, map[string]*_bintree_t{
	}},
	"000006_sharing_table.down.sql": &_bintree_t{_000006_sharing_table_down_sql, map[string]*_bintree_t{
	}},
	"000006_sharing_table.up.sql": &_bintree_t{_000006_sharing_table_up_sql, map[string]*_bintree_t{
	}},
	"000007_workspaces_table.down.sql": &_bintree_t{_000007_workspaces_table_down_sql, map[string]*_bintree_t{
	}},
	"000007_workspaces_table.up.sql": &_bintree_t{_000007_workspaces_table_up_sql, map[string]*_bintree_t{
	}},
	"000008_teams.down.sql": &_bintree_t{_000008_teams_down_sql, map[string]*_bintree_t{
	}},
	"000008_teams.up.sql": &_bintree_t{_000008_teams_up_sql, map[string]*_bintree_t{
	}},
	"000009_blocks_history.down.sql": &_bintree_t{_000009_blocks_history_down_sql, map[string]*_bintree_t{
	}},
	"000009_blocks_history.up.sql": &_bintree_t{_000009_blocks_history_up_sql, map[string]*_bintree_t{
	}},
	"000010_blocks_created_by.down.sql": &_bintree_t{_000010_blocks_created_by_down_sql, map[string]*_bintree_t{
	}},
	"000010_blocks_created_by.up.sql": &_bintree_t{_000010_blocks_created_by_up_sql, map[string]*_bintree_t{
	}},
	"000011_match_collation.down.sql": &_bintree_t{_000011_match_collation_down_sql, map[string]*_bintree_t{
	}},
	"000011_match_collation.up.sql": &_bintree_t{_000011_match_collation_up_sql, map[string]*_bintree_t{
	}},
}}
