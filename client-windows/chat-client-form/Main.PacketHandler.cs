using chat_client_form.Net;
using net.msg;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

using Message = net.msg.Message;

namespace chat_client_form
{
	public partial class Main : Form
	{
		private void InitializeHandler()
		{
			ClientSocket.Register(net.msg.Message.Types.Type.Way, HandleWay);
		}

		private void HandleWay(ArraySegment<byte> body)
		{
			Console.WriteLine("HandleWay");

			var reqWay = new ReqWay()
			{
				Name = "testname",
				Email = "sample@email.com"
			};

			ClientSocket.Send(Message.Types.Type.ReqWay, reqWay);
		}

	}
}
