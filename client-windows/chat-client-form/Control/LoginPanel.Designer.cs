﻿namespace chat_client_form.Control
{
	partial class LoginPanel
	{
		/// <summary> 
		/// 필수 디자이너 변수입니다.
		/// </summary>
		private System.ComponentModel.IContainer components = null;

		/// <summary> 
		/// 사용 중인 모든 리소스를 정리합니다.
		/// </summary>
		/// <param name="disposing">관리되는 리소스를 삭제해야 하면 true이고, 그렇지 않으면 false입니다.</param>
		protected override void Dispose(bool disposing)
		{
			if (disposing && (components != null))
			{
				components.Dispose();
			}
			base.Dispose(disposing);
		}

		#region 구성 요소 디자이너에서 생성한 코드

		/// <summary> 
		/// 디자이너 지원에 필요한 메서드입니다. 
		/// 이 메서드의 내용을 코드 편집기로 수정하지 마세요.
		/// </summary>
		private void InitializeComponent()
		{
			this.textBox_Nickname = new System.Windows.Forms.TextBox();
			this.button_Login = new System.Windows.Forms.Button();
			this.SuspendLayout();
			// 
			// textBox_Nickname
			// 
			this.textBox_Nickname.Anchor = ((System.Windows.Forms.AnchorStyles)((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom)));
			this.textBox_Nickname.Location = new System.Drawing.Point(215, 118);
			this.textBox_Nickname.Name = "textBox_Nickname";
			this.textBox_Nickname.Size = new System.Drawing.Size(100, 21);
			this.textBox_Nickname.TabIndex = 0;
			// 
			// button_Login
			// 
			this.button_Login.Anchor = System.Windows.Forms.AnchorStyles.Top;
			this.button_Login.Location = new System.Drawing.Point(215, 146);
			this.button_Login.Name = "button_Login";
			this.button_Login.Size = new System.Drawing.Size(100, 23);
			this.button_Login.TabIndex = 1;
			this.button_Login.Text = "Login";
			this.button_Login.UseVisualStyleBackColor = true;
			this.button_Login.Click += new System.EventHandler(this.Button_Login_Click);
			// 
			// LoginPanel
			// 
			this.AutoScaleDimensions = new System.Drawing.SizeF(7F, 12F);
			this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
			this.Controls.Add(this.button_Login);
			this.Controls.Add(this.textBox_Nickname);
			this.Name = "LoginPanel";
			this.Size = new System.Drawing.Size(580, 353);
			this.Load += new System.EventHandler(this.LoginPanel_Load);
			this.ResumeLayout(false);
			this.PerformLayout();

		}

		#endregion

		private System.Windows.Forms.TextBox textBox_Nickname;
		private System.Windows.Forms.Button button_Login;
	}
}
