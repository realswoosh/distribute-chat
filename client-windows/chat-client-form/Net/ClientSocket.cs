using Google.Protobuf;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using net.msg;

namespace chat_client_form.Net
{
	public sealed class ClientSocket
	{
		private static volatile ClientSocket instance;
		private static object syncRoot = new Object();

		private Socket socket;
		private const int BufferSize = 1024;
		private byte[] buffer = new byte[BufferSize];
		private byte[] received = new byte[BufferSize * 8];
		private int receiveSize = 0;

		private static readonly ManualResetEvent DoneConnect = new ManualResetEvent(false);
		private static readonly ManualResetEvent DoneSend = new ManualResetEvent(false);
		private static readonly ManualResetEvent DoneReceive = new ManualResetEvent(false);

		private Dictionary<Message.Types.Type, Action<IMessage>> msgHandlerDictionary = new Dictionary<Message.Types.Type, Action<IMessage>>();

		private ClientSocket()
		{
			socket = new Socket(SocketType.Stream, ProtocolType.Tcp);
		}

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
			IPEndPoint remoteEp = new IPEndPoint(IPAddress.Parse(addr), port);
			socket.BeginConnect(remoteEp,
				new AsyncCallback(ConnectCallback), socket);
			DoneConnect.WaitOne();
		}

		private void ConnectCallback(IAsyncResult ar)
		{
			try
			{
				Socket client = (Socket)ar.AsyncState;
				client.EndConnect(ar);
				DoneConnect.Set();

				Console.WriteLine("Connect complete!!");

				Receive();
			}
			catch (Exception e)
			{

			}
		}

		private void Receive()
		{
			try
			{
				socket.BeginReceive(buffer, 0, BufferSize, 0,
					new AsyncCallback(ReceiveCallback), socket);
			}
			catch (Exception e)
			{

			}
		}

		public T ToObject<T>(byte[] buf) where T : IMessage<T>, new()
		{
			if (buf == null)
				return default(T);

			using (MemoryStream ms = new MemoryStream())
			{
				ms.Write(buf, 0, buf.Length);
				ms.Seek(0, SeekOrigin.Begin);

				MessageParser<T> parser = new MessageParser<T>(() => new T());
				return parser.ParseFrom(ms);
			}
		}

		private void ReceiveCallback(IAsyncResult ar)
		{
			try
			{
				Socket socket = (Socket)ar.AsyncState;
				int bytesRead = socket.EndReceive(ar);

				if (bytesRead > 0)
				{
					using (var writer = new MemoryStream(received))
					{
						writer.Write(buffer, receiveSize, bytesRead);
					}

					receiveSize += bytesRead;

					Int32 readOffset = 0;

					while (Def.Constants.HeaderSize <= receiveSize - readOffset)
					{
						int packetSize = BitConverter.ToInt32(received, readOffset);
						if (packetSize >= receiveSize - readOffset)
						{
							byte[] packet = new byte[packetSize];

							using (var reader = new MemoryStream(received))
							{
								reader.Read(packet, 0, packetSize);
							}

							ArraySegment<byte> segMsg = new ArraySegment<byte>(packet, 4, 4);
							ArraySegment<byte> segBody = default;

							if (packetSize > Def.Constants.HeaderSize)
							{
								segBody = new ArraySegment<byte>(packet, Def.Constants.HeaderSize, packetSize - Def.Constants.HeaderSize);
							}

							PacketHandler(segMsg, segBody);

							readOffset += packetSize;
						}
						else
						{
							break;
						}
					}

					Console.WriteLine("receive {0} bytes", bytesRead);
				}
			}
			catch (Exception e)
			{

			}
		}

		private void PacketHandler(ArraySegment<byte> msg, ArraySegment<byte> body)
		{
			Int32 tmpMsg = BitConverter.ToInt32(msg.Array, msg.Offset);

			var msgType = (net.msg.Message.Types.Type)Enum.ToObject(typeof(net.msg.Message.Types.Type), tmpMsg);

			Console.WriteLine("msgType = {0}", msgType.ToString());
		}
	}
}
