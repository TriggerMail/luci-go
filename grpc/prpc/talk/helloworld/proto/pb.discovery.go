// AUTOGENERATED. DO NOT EDIT.

package helloworld

import discovery "github.com/TriggerMail/luci-go/grpc/discovery"

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

func init() {
	discovery.RegisterDescriptorSetCompressed(
		[]string{
			"helloworld.Greeter",
		},
		[]byte{31, 139,
			8, 0, 0, 0, 0, 0, 0, 255, 108, 147, 193, 79, 219, 72,
			20, 198, 61, 158, 73, 72, 94, 194, 134, 204, 18, 146, 141, 86,
			218, 183, 57, 0, 90, 177, 14, 100, 23, 237, 109, 37, 132, 80,
			75, 69, 57, 64, 233, 221, 56, 47, 246, 168, 142, 199, 157, 177,
			105, 243, 247, 244, 214, 67, 111, 253, 255, 170, 113, 98, 200, 161,
			23, 203, 223, 243, 123, 159, 127, 254, 230, 25, 190, 236, 192, 85,
			172, 131, 40, 49, 122, 169, 202, 101, 160, 77, 60, 77, 203, 72,
			77, 99, 147, 71, 211, 220, 93, 138, 48, 253, 48, 77, 40, 77,
			245, 39, 109, 210, 249, 52, 55, 186, 208, 91, 133, 160, 42, 72,
			120, 169, 76, 38, 208, 125, 237, 212, 29, 125, 44, 201, 22, 82,
			130, 200, 194, 37, 141, 24, 178, 227, 246, 93, 117, 63, 57, 4,
			216, 244, 228, 233, 74, 142, 96, 103, 73, 214, 134, 113, 221, 84,
			203, 217, 53, 236, 188, 50, 68, 5, 25, 249, 63, 180, 238, 195,
			85, 53, 37, 71, 193, 22, 193, 246, 203, 198, 7, 63, 121, 146,
			167, 171, 137, 247, 230, 91, 3, 154, 82, 8, 111, 196, 224, 59,
			3, 214, 149, 92, 120, 114, 246, 149, 225, 165, 206, 87, 70, 197,
			73, 129, 179, 211, 179, 115, 124, 151, 16, 222, 60, 92, 94, 227,
			69, 89, 36, 218, 216, 0, 47, 210, 20, 171, 6, 139, 134, 44,
			153, 39, 154, 7, 128, 15, 150, 80, 47, 176, 72, 148, 69, 171,
			75, 19, 17, 70, 122, 78, 168, 44, 198, 250, 137, 76, 70, 115,
			44, 179, 57, 25, 44, 18, 194, 139, 60, 140, 156, 177, 138, 40,
			179, 116, 130, 239, 201, 88, 165, 51, 156, 5, 167, 128, 69, 18,
			22, 24, 133, 25, 62, 18, 46, 116, 153, 205, 81, 101, 213, 212,
			205, 245, 229, 213, 237, 253, 21, 46, 84, 74, 1, 64, 11, 152,
			47, 121, 179, 37, 33, 0, 191, 233, 73, 209, 246, 186, 108, 60,
			169, 144, 99, 23, 147, 202, 98, 116, 128, 42, 34, 156, 211, 66,
			101, 170, 80, 58, 11, 0, 0, 120, 211, 99, 146, 183, 91, 61,
			56, 4, 209, 244, 124, 79, 242, 142, 127, 62, 254, 13, 235, 80,
			215, 14, 22, 67, 44, 45, 153, 0, 160, 11, 13, 215, 199, 36,
			239, 52, 127, 169, 149, 47, 121, 103, 239, 247, 90, 113, 201, 59,
			71, 103, 240, 47, 248, 194, 147, 162, 231, 73, 54, 62, 198, 237,
			243, 192, 72, 103, 69, 168, 50, 91, 125, 142, 51, 62, 178, 232,
			54, 96, 205, 36, 156, 123, 175, 181, 15, 255, 129, 16, 21, 83,
			223, 151, 227, 191, 240, 54, 92, 86, 65, 186, 33, 215, 189, 14,
			154, 48, 39, 99, 117, 134, 197, 6, 54, 0, 216, 133, 134, 27,
			20, 82, 244, 253, 222, 129, 227, 114, 178, 225, 140, 90, 181, 98,
			146, 247, 219, 187, 181, 226, 146, 247, 247, 250, 240, 55, 248, 130,
			73, 49, 240, 70, 108, 252, 39, 190, 108, 202, 11, 113, 248, 156,
			233, 6, 150, 49, 201, 7, 45, 9, 39, 32, 4, 115, 176, 67,
			127, 48, 254, 3, 223, 174, 119, 181, 230, 125, 62, 136, 130, 62,
			215, 132, 172, 34, 28, 250, 131, 253, 138, 130, 85, 132, 195, 13,
			33, 171, 8, 135, 237, 189, 90, 113, 201, 135, 191, 238, 63, 54,
			171, 95, 235, 159, 31, 1, 0, 0, 255, 255, 206, 118, 48, 11,
			166, 3, 0, 0},
	)
}

// FileDescriptorSet returns a descriptor set for this proto package, which
// includes all defined services, and all transitive dependencies.
//
// Will not return nil.
//
// Do NOT modify the returned descriptor.
func FileDescriptorSet() *descriptor.FileDescriptorSet {
	// We just need ONE of the service names to look up the FileDescriptorSet.
	ret, err := discovery.GetDescriptorSet("helloworld.Greeter")
	if err != nil {
		panic(err)
	}
	return ret
}
