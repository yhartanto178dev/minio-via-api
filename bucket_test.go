package main

import "testing"

func TestMakeBucket(t *testing.T) {
	Success, _ := MakeBucket("arch-bucket", "192.168.75.1:9000", "uqrdCSsxK1MbnwJyhEoG", "QbBetsiMmGFbFJrgyTjafpMHxzYcERq06KiD0tFY")
	if Success != "Success" {
		t.Errorf("MakeBucket() = %s; want Success", Success)
	}
}

func TestListBuckets(t *testing.T) {
	buckets, _ := ListBuckets("192.168.75.1:9000", "uqrdCSsxK1MbnwJyhEoG", "QbBetsiMmGFbFJrgyTjafpMHxzYcERq06KiD0tFY")
	if len(buckets) == 0 {
		t.Errorf("ListBuckets() = %d; want > 0", len(buckets))
	}
	for _, bucket := range buckets {
		if bucket.Name != "arch-bucket" {
			t.Errorf("ListBuckets() = %s; want arch-bucket", bucket.Name)
		}
	}
}
