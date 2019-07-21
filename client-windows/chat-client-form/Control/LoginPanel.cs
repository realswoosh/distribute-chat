using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Drawing;
using System.Data;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace chat_client_form.Control
{
	public partial class LoginPanel : UserControl
	{
		public event EventHandler ButtonLoginClick;

		public LoginPanel()
		{
			InitializeComponent();
		}

		private void LoginPanel_Load(object sender, EventArgs e)
		{
			textBox_Nickname.Left = (this.ClientSize.Width - textBox_Nickname.Width) / 2;
			
			button_Login.Left = (this.ClientSize.Width - button_Login.Width) / 2;
			button_Login.Top = textBox_Nickname.Top + textBox_Nickname.Height + 5;
		}

		private void Button_Login_Click(object sender, EventArgs e)
		{
			if (this.ButtonLoginClick != null)
				this.ButtonLoginClick(this, new EventArgs());
		}
	}
}
