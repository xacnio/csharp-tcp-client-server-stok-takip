using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace WindowsFormsApp1
{
    internal static class Program
    {
        /// <summary>
        /// Uygulamanın ana girdi noktası.
        /// </summary>
        [STAThread]
        static void Main()
        {
            Application.EnableVisualStyles();
            Application.SetCompatibleTextRenderingDefault(false);

            CSocket.Connect("127.0.0.1", 4444);
            if (!CSocket.isConnected)
            {
                MessageBox.Show(CSocket.sException.Message, "Bağlantı Hatası");
            }
            else
            {
                Application.Run(new Form1());
            }
        }
    }
}
