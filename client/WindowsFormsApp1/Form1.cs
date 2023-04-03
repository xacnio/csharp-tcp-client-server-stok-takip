using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Net.Sockets;
using System.Net;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using WindowsFormsApp1.Properties;
using static System.Windows.Forms.VisualStyles.VisualStyleElement;
using System.Text.RegularExpressions;

namespace WindowsFormsApp1
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private bool focusedPrice = false;

        private int currentPage = 1;
        private int maxPage = 0;

        private void Form1_Load(object sender, EventArgs e)
        {
            GetList(1);
        }

        private void Form1_FormClosing(object sender, FormClosingEventArgs e)
        {
            CSocket.Disconnect();
        }

        private void GetList(int page)
        {
            nextButton.Enabled = false;
            backButton.Enabled = false;
            Cursor = Cursors.WaitCursor;
            
            string data = "";

            if(textBox1.Text != "")
            {
                data = CSocket.Send("search:" + page.ToString() + ":" + textBox1.Text);
            } 
            else
            {
                data = CSocket.Send("list:" + page.ToString());
            }

            Stoklar stoklar = Stoklar.FromJSON(data);

            Cursor = Cursors.Default;

            if (stoklar == null)
            {
                return;
            }

            dataGridView1.Rows.Clear();

            foreach (Stok var in stoklar.items)
            {
                object[] row = new object[] { var.id, var.isim, var.fiyat, var.sayi };
                dataGridView1.Rows.Add(row);
            }

            maxPage = stoklar.total / stoklar.max_item_count;
            if (stoklar.total % stoklar.max_item_count > 0) maxPage++;
            if (currentPage == 1 && maxPage == 0) currentPage = 0;
            label1.Text = currentPage + "/" + maxPage;
            label1.Visible = true;

            if(stoklar.total > 0)
            {
                label6.Text = "Toplam: " + stoklar.total.ToString();
                label6.Visible = true;
            }
            else
            {
                label6.Visible = false;
            }


            if (stoklar.total > stoklar.max_item_count)
            {
                if (currentPage < maxPage) nextButton.Enabled = true;
                if (currentPage > 1) backButton.Enabled = true;
            }
        }

        private void backButton_Click(object sender, EventArgs e)
        {
            GetList(--currentPage);
        }

        private void nextButton_Click(object sender, EventArgs e)
        {
            GetList(++currentPage);
        }



        private void dataGridView1_UserDeletingRow(object sender, DataGridViewRowCancelEventArgs e)
        {
            string id = e.Row.Cells[0].Value.ToString();
            String data = CSocket.Send("delete:" + id);
            if(data != "OK")
            {
                MessageBox.Show("Hata: " + data, "Silinemedi", MessageBoxButtons.OK, MessageBoxIcon.Error);
                e.Cancel = true;
            }
        }

        private void dataGridView1_UserDeletedRow(object sender, DataGridViewRowEventArgs e)
        {
            GetList(currentPage);
        }

        private void textBox1_TextChanged(object sender, EventArgs e)
        {
            currentPage = 1;
            GetList(currentPage);
        }

        private void dataGridView1_SelectionChanged(object sender, EventArgs e)
        {
            // current selected row
            DataGridViewRow r = dataGridView1.CurrentRow;
            if(r == null)
            {
                productName.Text = "";
                productPrice.Text = "";
                productCount.Text = "";
                return;
            }
            productName.Text = r.Cells[1].Value.ToString();
            productPrice.Text = r.Cells[2].Value.ToString();
            productCount.Text = r.Cells[3].Value.ToString();
        }

        private void silToolStripMenuItem_Click(object sender, EventArgs e)
        {
            int r = dataGridView1.Rows.GetFirstRow(DataGridViewElementStates.Selected);
            String data = CSocket.Send("delete:" + dataGridView1.Rows[r].Cells[0].Value.ToString());
            if (data != "OK")
            {
                MessageBox.Show("Hata: " + data, "Silinemedi", MessageBoxButtons.OK, MessageBoxIcon.Error);
            }
            else
            {
                GetList(currentPage);
            }
        }

        private void dataGridView1_MouseDown(object sender, MouseEventArgs e)
        {
            if (e.Button == MouseButtons.Right)
            {
                var hti = dataGridView1.HitTest(e.X, e.Y);
                dataGridView1.ClearSelection();
                if(hti.RowIndex >= 0)
                    dataGridView1.Rows[hti.RowIndex].Selected = true;
            }
        }

        private void tümünüSilToolStripMenuItem_Click(object sender, EventArgs e)
        {
            // Confirm dialog
            DialogResult confirm = MessageBox.Show("Silmek istediğinize emin misiniz?", "Onay", MessageBoxButtons.YesNo, MessageBoxIcon.Warning);
            if(confirm == DialogResult.Yes)
            {
                DataGridViewRowCollection copyDataViewRows = dataGridView1.Rows;
                foreach (DataGridViewRow row in dataGridView1.Rows)
                {
                    string id = row.Cells[0].Value.ToString();
                    string data = CSocket.Send("delete:" + id);
                    if (data != "OK")
                    {
                        copyDataViewRows.RemoveAt(row.Index);
                    }
                }
                dataGridView1.Rows.Clear();
                foreach(DataGridViewRow row in copyDataViewRows)
                {
                    dataGridView1.Rows.Add(row.Cells);
                }
                GetList(currentPage);
            }
        }

        private void productPrice_TextChanged(object sender, EventArgs e)
        {
            if (focusedPrice) return;

            Regex number = new Regex(@"^[0-9]+$");
            if(number.IsMatch(productPrice.Text))
            {
                decimal price = decimal.Parse(productPrice.Text);
                productPrice.Text = String.Format("{0:c2}", price);
            }
        }

        private void productPrice_Enter(object sender, EventArgs e)
        {
            if(!focusedPrice)
            {
                int r = productPrice.Text.IndexOf(",");
                if (r != -1 && !focusedPrice)
                {
                    productPrice.Select(r, 0);
                }
                focusedPrice = true;
            }
            productPrice.Text = Regex.Replace(productPrice.Text, @"[^0-9,]+", "");
        }

        private void productPrice_Leave(object sender, EventArgs e)
        {
            focusedPrice = false;
            Regex number = new Regex(@"^[0-9,]+$");
            if (number.IsMatch(productPrice.Text))
            {
                decimal price = decimal.Parse(productPrice.Text);
                productPrice.Text = String.Format("{0:c2}", price);
            }
        }

        private void button1_Click(object sender, EventArgs e)
        {
            string price = Regex.Replace(productPrice.Text, @"[^0-9,]+", "");
            decimal parsedPrice = decimal.Parse(price);
            price = Regex.Replace(price, "[,]", ".");

            string name = productName.Text;
            string count = productCount.Text;

            string sendData = String.Format("upsert:{0}:{1}:{2}", name, price, count);
            string data = CSocket.Send(sendData);
            if (data == "OK")
            {
                saveMessage.Text = "Kaydedildi!";
                saveMessage.Visible = true;
                int r = dataGridView1.Rows.GetFirstRow(DataGridViewElementStates.Selected);
                if (r != -1 && dataGridView1.Rows[r].Cells[1].Value.Equals(name))
                {
                    dataGridView1.Rows[r].Cells[2].Value = parsedPrice;
                    dataGridView1.Rows[r].Cells[3].Value = int.Parse(count);
                }
                else
                {
                    currentPage = 1;
                    GetList(currentPage);
                }
            }
            else
            {
                saveMessage.Text = data;
                saveMessage.Visible = true;
            }
        }

        private void button3_Click(object sender, EventArgs e)
        {
            try
            {
                int count = Int32.Parse(productCount.Text);
                productCount.Text = (++count).ToString();
            }
            catch(Exception ex) 
            {
                productCount.Text = "0";
            }
            
        }

        private void button2_Click(object sender, EventArgs e)
        {
            try
            {
                int count = Int32.Parse(productCount.Text);
                if(count > 0) productCount.Text = (--count).ToString();
            }
            catch (Exception ex) 
            {
                productCount.Text = "0";
            }
        }

        private void button4_Click(object sender, EventArgs e)
        {
            dataGridView1.CurrentCell = null;
            productName.Text = "";
            productCount.Text = "0";
            productPrice.Text = "0";
        }

        private void productPrice_KeyDown(object sender, KeyEventArgs e)
        {

        }

        private void productPrice_KeyPress(object sender, KeyPressEventArgs e)
        {
            if (!char.IsControl(e.KeyChar) && !char.IsDigit(e.KeyChar) && (e.KeyChar != ','))
            {
                e.Handled = true;
            }
            if ((e.KeyChar == ',') && ((sender as System.Windows.Forms.TextBox).Text.IndexOf(',') > -1))
            {
                e.Handled = true;
            }
        }

        private void productPrice_GotFocus(object sender, EventArgs e)
        {
            int r = productPrice.Text.IndexOf(",");
            if(r != -1)
            {
                productPrice.Select(r, 0);
            }
        }

        private void productCount_KeyPress(object sender, KeyPressEventArgs e)
        {
            if (!char.IsControl(e.KeyChar) && !char.IsDigit(e.KeyChar))
            {
                e.Handled = true;
            }
        }

        private void productPrice_Click(object sender, EventArgs e)
        {
            label1.Focus();
            productPrice.Focus();
        }

        private void productPrice_MouseClick(object sender, MouseEventArgs e)
        {

        }
    }
}
