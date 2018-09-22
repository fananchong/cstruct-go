#include <string>

#pragma pack(1)

struct StructA {
	uint8_t A1;
	uint32_t A2;
	uint8_t A3[5];
};


struct StructB {
	uint8_t B1;
	StructA B2;
	uint16_t B3;
	float B4;
	StructA B5[3];
};

int main()
{
	StructB b;
	b.B1 = 127;
	b.B2.A1 = 56;
	b.B2.A2 = 999;
	b.B2.A3[0] = 0;
	b.B2.A3[1] = 1;
	b.B2.A3[2] = 2;
	b.B2.A3[3] = 3;
	b.B2.A3[4] = 4;
	b.B3 = 8888;
	b.B4 = 88.8f;
	b.B5[0] = b.B2;
	b.B5[1] = b.B2;
	b.B5[2] = b.B2;

	printf("len(b) = %llu\n", sizeof(b));
	printf("struct data len = %llu\n", sizeof(b));
	printf("struct data is:\n");

	unsigned char buff[1024];
	memcpy(buff, &b, sizeof(b));
	for (int i = 0; i < sizeof(b); i++) {
		printf("%d ", buff[i]);
	}
	return 0;
}
