package main

import "testing"

func TestMakeBucket(t *testing.T) {
	Success, _ := MakeBucket("arch-buck", "192.168.75.1:9000", "ypeUmgEkLH6Q5nyZOH9c", "qY1570kfxvqGAozRopSE5d1rkujiU95Uk7JETlKo")
	if Success != "Success" {
		t.Errorf("MakeBucket() = %s; want Success", Success)
	}
}

func TestListBuckets(t *testing.T) {
	buckets, _ := ListBuckets("192.168.75.1:9000", "ypeUmgEkLH6Q5nyZOH9c", "qY1570kfxvqGAozRopSE5d1rkujiU95Uk7JETlKo")
	if len(buckets) == 0 {
		t.Errorf("ListBuckets() = %d; want > 0", len(buckets))
	}
	for _, bucket := range buckets {
		t.Logf("Bucket name: %s\n", bucket.Name)
		t.Logf("Bucket created: %s\n", bucket.CreationDate)
	}
}

func TestFPutObject(t *testing.T) {
	Success, _ := FPutObjects("archive-buck", "example/BILIONAIRE.mp4", "example/BILIONAIRE.mp4", "192.168.75.1:9000", "ypeUmgEkLH6Q5nyZOH9c", "qY1570kfxvqGAozRopSE5d1rkujiU95Uk7JETlKo")
	if Success != "Success" {
		t.Errorf("FPutObjects() = %s; want Success", Success)
	}
}
func TestFGetObject(t *testing.T) {
	Success, _ := FGetObjects("archive-buck", "example/31-01-2025-PERMASALAHAN-KORWILCAM-SUMPIUH.png", "example/31-01-2025-PERMASALAHAN-KORWILCAM-SUMPIUH.png", "192.168.75.1:9000", "ypeUmgEkLH6Q5nyZOH9c", "qY1570kfxvqGAozRopSE5d1rkujiU95Uk7JETlKo")
	if Success != "Success" {
		t.Errorf("FGetObjects() = %s; want Success", Success)
	}
}
func TestListObjects(t *testing.T) {
	objects, _ := ListObjects("archive-buck", "192.168.75.1:9000", "ypeUmgEkLH6Q5nyZOH9c", "qY1570kfxvqGAozRopSE5d1rkujiU95Uk7JETlKo", "example/")
	if len(objects) == 0 {
		t.Errorf("ListObjects() = %d; want > 0", len(objects))
	}
	for _, object := range objects {
		t.Logf("Object name: %s\n", object.Key)
		sizeInMB := float64(object.Size) / (1024 * 1024)
		t.Logf("Object size: %.2f MB\n", sizeInMB)
	}
}
