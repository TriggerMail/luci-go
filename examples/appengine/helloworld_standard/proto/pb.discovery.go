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
			8, 0, 0, 0, 0, 0, 0, 255, 116, 147, 193, 110, 219, 56,
			16, 134, 69, 145, 81, 236, 177, 179, 118, 184, 182, 17, 24, 11,
			236, 192, 135, 205, 30, 22, 114, 226, 69, 123, 41, 90, 32, 8,
			130, 54, 69, 208, 67, 210, 244, 90, 40, 210, 88, 34, 32, 147,
			42, 41, 165, 245, 43, 245, 222, 7, 43, 250, 2, 133, 24, 171,
			205, 161, 189, 233, 227, 12, 249, 127, 212, 72, 240, 45, 130, 235,
			220, 196, 105, 97, 205, 70, 53, 155, 216, 216, 124, 89, 54, 169,
			90, 210, 167, 100, 83, 149, 228, 150, 73, 85, 145, 206, 149, 166,
			101, 65, 101, 105, 62, 26, 91, 102, 239, 93, 157, 232, 44, 177,
			217, 178, 178, 166, 54, 143, 42, 177, 95, 144, 240, 115, 101, 177,
			128, 225, 171, 150, 174, 233, 67, 67, 174, 150, 18, 132, 78, 54,
			116, 196, 144, 253, 219, 191, 246, 207, 139, 127, 0, 118, 61, 85,
			185, 149, 71, 176, 191, 33, 231, 146, 188, 107, 234, 112, 117, 9,
			251, 47, 45, 81, 77, 86, 190, 128, 222, 77, 178, 245, 187, 228,
			81, 252, 200, 224, 113, 216, 124, 246, 139, 74, 85, 110, 23, 193,
			235, 175, 2, 34, 41, 68, 48, 99, 240, 133, 1, 27, 74, 46,
			2, 185, 250, 204, 240, 220, 84, 91, 171, 242, 162, 198, 213, 201,
			233, 83, 124, 91, 16, 94, 221, 158, 95, 226, 89, 83, 23, 198,
			186, 24, 207, 202, 18, 125, 131, 67, 75, 142, 236, 61, 101, 49,
			224, 173, 35, 52, 107, 172, 11, 229, 208, 153, 198, 166, 132, 169,
			201, 8, 149, 195, 220, 220, 147, 213, 148, 97, 163, 51, 178, 88,
			23, 132, 103, 85, 146, 182, 7, 171, 148, 180, 163, 255, 240, 29,
			89, 167, 140, 198, 85, 124, 2, 88, 23, 73, 141, 105, 162, 241,
			142, 112, 109, 26, 157, 161, 210, 126, 215, 213, 229, 249, 197, 155,
			155, 11, 92, 171, 146, 98, 128, 30, 176, 80, 242, 168, 39, 33,
			134, 48, 10, 164, 232, 7, 67, 54, 95, 120, 229, 188, 125, 77,
			74, 231, 216, 10, 170, 148, 48, 163, 181, 210, 170, 86, 70, 199,
			0, 0, 60, 10, 152, 228, 253, 222, 8, 254, 6, 17, 5, 97,
			32, 249, 32, 124, 50, 151, 120, 67, 58, 115, 152, 252, 56, 0,
			96, 8, 123, 109, 3, 147, 124, 16, 253, 209, 81, 40, 249, 96,
			252, 87, 71, 92, 242, 193, 241, 41, 60, 135, 80, 4, 82, 140,
			130, 67, 54, 63, 245, 26, 246, 97, 14, 184, 27, 32, 166, 70,
			215, 137, 210, 173, 88, 123, 161, 198, 145, 61, 118, 216, 126, 3,
			15, 86, 162, 141, 25, 245, 38, 48, 0, 33, 188, 213, 56, 148,
			112, 0, 123, 45, 8, 41, 198, 225, 104, 214, 134, 182, 184, 215,
			22, 123, 29, 49, 201, 199, 253, 131, 142, 184, 228, 227, 241, 33,
			60, 131, 80, 48, 41, 38, 193, 140, 205, 151, 59, 33, 87, 25,
			237, 232, 119, 70, 221, 181, 221, 131, 14, 99, 146, 79, 122, 210,
			235, 176, 86, 103, 26, 78, 189, 14, 243, 58, 211, 112, 50, 241,
			145, 204, 235, 76, 119, 58, 204, 235, 76, 251, 227, 142, 184, 228,
			211, 63, 39, 119, 145, 255, 59, 254, 255, 30, 0, 0, 255, 255,
			132, 239, 156, 21, 118, 3, 0, 0},
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
