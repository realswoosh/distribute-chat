using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace chat_client_form
{
	public partial class Main : Form
	{
		public Main()
		{
			InitializeComponent();

			InitializeHandler();
		}

		private void Main_Load(object sender, EventArgs e)
		{
			loginPanel.ButtonLoginClick += new EventHandler(UserPanel_LoginButtonClick);
		}
	}
}
