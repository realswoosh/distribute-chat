using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Google.Protobuf;
using net.msg;

namespace chat_client_form.Net
{
	public class Packet
	{
		private static readonly Int32 HeaderSize = 8;
		private Int32 size;
		private Message.Types.Type type;
		private byte[] data;

		public byte[] Buffer
		{
			get
			{
				byte[] tmp = new byte[HeaderSize + data.Length];
				using (MemoryStream ms = new MemoryStream(tmp))
				using (BinaryWriter bw = new BinaryWriter(ms))
				{
					bw.Write(size);
					bw.Write((int)type);
					bw.Write(data);
					return tmp;
				}
			}
		}

		public Packet()
		{

		}

		public Packet(Message.Types.Type type, IMessage data)
		{
			this.type = type;
			this.data = data.ToByteArray();

			size = HeaderSize + this.data.Length;
		}
	}
}
