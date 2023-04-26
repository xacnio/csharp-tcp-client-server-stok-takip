using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Runtime.Serialization;
using System.Runtime.Serialization.Json;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace WindowsFormsApp1
{
    [DataContract]
    internal class Stok
    {
        [DataMember]
        public int id { get; private set; }
        [DataMember]
        public string isim { get; private set; }
        [DataMember]
        public float fiyat { get; private set; }
        [DataMember]

        public float promosyonlu_fiyat { get; private set; }
        [DataMember]
        public float sayi { get; private set; }
    }

    [DataContract]
    internal class Stoklar
    {
        [DataMember]
        public int total { get; private set; }
        [DataMember]
        public int max_item_count { get; private set; }
        [DataMember]
        public Stok[] items { get; private set; }

        public static Stoklar FromJSON(string JSONdata)
        {
            try
            {
                DataContractJsonSerializer jsonSer = new DataContractJsonSerializer(typeof(Stoklar));
                MemoryStream stream = new MemoryStream(Encoding.UTF8.GetBytes(JSONdata));
                Stoklar objStoklar = (Stoklar)jsonSer.ReadObject(stream);
                return objStoklar;
            }
            catch(Exception e)
            {
                MessageBox.Show(JSONdata, "Ayrıştırma Hatası");
                Console.WriteLine(e.Message);
            }
            return null;
        }
    }
}
