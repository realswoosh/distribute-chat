using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Sockets;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace chat_client_form.Net
{
	public sealed class ClientSocket
	{
		private static volatile ClientSocket instance;
		private static object syncRoot = new Object();

		private Socket socket;
		private const int BufferSize = 1024;
		private byte[] buffer = new byte[BufferSize];

		private static readonly ManualResetEvent DoneConnect = new ManualResetEvent(false);
		private static readonly ManualResetEvent DoneSend = new ManualResetEvent(false);
		private static readonly ManualResetEvent DoneReceive = new ManualResetEvent(false);

		private ClientSocket() { }

		public static ClientSocket Instance
		{
			get
			{
				if (instance == null)
				{
					lock (syncRoot)
					{
						if (instance == null)
							instance = new ClientSocket();
					}
				}

				return instance;
			}
		}

		public void Dial(string addr, int port)
		{

		}
	}
}
