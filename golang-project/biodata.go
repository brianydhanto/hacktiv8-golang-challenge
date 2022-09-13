package main

import (
    "fmt"
    "os"
)

type Person struct {
  id int
  nama string
  alamat string
  pekerjaan string
  alasan string
}


func main() {
   var person1 Person;
   var person2 Person;
   var person3 Person;
   var indexTemp int;

   sliceOfName := []string{"Brian", "Daniel", "Mawar"};

   person1.id = 1;
   person1.nama = "Brian";
   person1.alamat = "Jalan Bangka Raya 1";
   person1.pekerjaan = "Frontend Enginner";
   person1.alasan = "Alasan Brian";

   person2.id = 2;
   person2.nama = "Daniel";
   person2.alamat = "Jalan Bangka Raya 2";
   person2.pekerjaan = "Backend Enginner";
   person2.alasan = "Alasan Daniel";

   person3.id = 3;
   person3.nama = "Mawar";
   person3.alamat = "Jalan Bangka Raya 3";
   person3.pekerjaan = "Data Enginner";
   person3.alasan = "Alasan Mawar";

   sliceOfPerson := []Person{person1, person2, person3};

   for idx, item :=range sliceOfName {
        if item == os.Args[1] {
            indexTemp = idx;
        }
   }

   fmt.Printf("ID: %d\n", sliceOfPerson[indexTemp].id);
   fmt.Printf("Nama: %s\n", sliceOfPerson[indexTemp].nama);
   fmt.Printf("Alamat: %s\n", sliceOfPerson[indexTemp].alamat);
   fmt.Printf("Pekerjaan: %s\n", sliceOfPerson[indexTemp].pekerjaan);
   fmt.Printf("Alasan: %s\n", sliceOfPerson[indexTemp].alasan);
}