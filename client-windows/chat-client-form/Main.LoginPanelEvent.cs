﻿using chat_client_form.Net;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace chat_client_form
{
	public partial class Main : Form
	{
		public void UserPanel_LoginButtonClick(object sender, EventArgs e)
		{
			Console.WriteLine("UserPanel_LoginButtonClick");
			//your code here
			ClientSocket.Instance.Dial("127.0.0.1", 20055);

		}
	}
}
