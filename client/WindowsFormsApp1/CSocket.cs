using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Sockets;
using System.Net;
using System.Text;
using System.Threading.Tasks;

namespace WindowsFormsApp1
{
    internal class CSocket
    { 
        const int MAX_READ_SIZE_SOCKET = 5 * 1024 * 1024; // 5 MB
        private static Socket sender;
        public static bool isConnected { get; private set; }
        public static Exception sException { get; private set; }

        public static String Send(String message)
        {
            if(!isConnected) return "Bağlantı sağlanamadı";

            byte[] bytes = new byte[MAX_READ_SIZE_SOCKET];
            byte[] msg = Encoding.UTF8.GetBytes(message);
            try
            {
                int bytesSent = sender.Send(msg);

                int bytesRec = sender.Receive(bytes);

                return Encoding.UTF8.GetString(bytes, 0, bytesRec);

            }
            catch (ArgumentNullException ane)
            {
                return ane.ToString();
            }
            catch (SocketException se)
            {
                isConnected = false;
                sException = se;
                return se.ToString();
            }
            catch (Exception e)
            {
                return e.ToString();
            }
        }

        public static void Connect(String ip, int port)
        { 
            try
            {
                IPHostEntry host = Dns.GetHostEntry("127.0.0.1");
                IPAddress ipAddress = host.AddressList[0];
                IPEndPoint remoteEP = new IPEndPoint(ipAddress, 4444);

                sender = new Socket(ipAddress.AddressFamily,
                    SocketType.Stream, ProtocolType.Tcp);

                try
                {
                    sender.Connect(remoteEP);
                    isConnected = true;    
                }
                catch (ArgumentNullException ane)
                {
                    isConnected = false;
                    sException = ane;
                }
                catch (SocketException se)
                {
                    isConnected = false;
                    sException = se;
                }
                catch (Exception e)
                {
                    isConnected = false;
                    sException = e;
                }

            }
            catch (Exception e)
            {
                isConnected = false;
                sException = e;
            }
        }

        public static void Disconnect()
        {
            if (isConnected)
            {
                sender.Disconnect(false);
                isConnected = false;
            }
        }
    }
}
