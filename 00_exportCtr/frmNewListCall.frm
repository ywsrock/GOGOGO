VERSION 5.00
Object = "{831FDD16-0C5C-11D2-A9FC-0000F8754DA1}#2.2#0"; "MSCOMCTL.OCX"
Object = "{F9043C88-F6F2-101A-A3C9-08002B2F49FB}#1.2#0"; "COMDLG32.OCX"
Object = "{42FFC9A8-7FC5-4D28-95E0-9273BAC25E7D}#3.0#0"; "lvwODDB2.ocx"
Begin VB.Form frmNewListCall 
   BackColor       =   &H00008000&
   BorderStyle     =   1  '�Œ�(����)
   Caption         =   "�˓d���X�g�ꗗ"
   ClientHeight    =   12030
   ClientLeft      =   45
   ClientTop       =   1440
   ClientWidth     =   17700
   ControlBox      =   0   'False
   BeginProperty Font 
      Name            =   "�l�r �o����"
      Size            =   11.25
      Charset         =   128
      Weight          =   700
      Underline       =   0   'False
      Italic          =   0   'False
      Strikethrough   =   0   'False
   EndProperty
   Icon            =   "frmNewListCall.frx":0000
   LinkTopic       =   "Form2"
   MaxButton       =   0   'False
   ScaleHeight     =   12030
   ScaleWidth      =   17700
   Begin VB.PictureBox Picture1 
      BackColor       =   &H00C0E0FF&
      Height          =   12015
      Left            =   0
      ScaleHeight     =   11955
      ScaleWidth      =   17715
      TabIndex        =   0
      Top             =   0
      Width           =   17775
      Begin lvwODDB2.ucLvwODDB2 lvwCallList 
         Height          =   4575
         Left            =   240
         TabIndex        =   90
         Top             =   4440
         Width           =   17235
         _ExtentX        =   30401
         _ExtentY        =   8070
      End
      Begin VB.Frame Frame3 
         BackColor       =   &H00C0E0FF&
         Caption         =   "���s���[�h"
         Height          =   2055
         Left            =   14640
         TabIndex        =   83
         Top             =   1800
         Width           =   2895
         Begin VB.PictureBox Picture2 
            BackColor       =   &H0080C0FF&
            Height          =   345
            Left            =   360
            ScaleHeight     =   285
            ScaleWidth      =   2235
            TabIndex        =   84
            Top             =   360
            Width           =   2295
            Begin VB.OptionButton optMod 
               BackColor       =   &H0080C0FF&
               Caption         =   "AND"
               Height          =   255
               Index           =   1
               Left            =   1320
               TabIndex        =   55
               Top             =   30
               Value           =   -1  'True
               Width           =   915
            End
            Begin VB.OptionButton optMod 
               BackColor       =   &H0080C0FF&
               Caption         =   "OR"
               Height          =   255
               Index           =   0
               Left            =   180
               TabIndex        =   54
               Top             =   30
               Width           =   915
            End
         End
         Begin VB.CommandButton btnSearch 
            Caption         =   "����"
            BeginProperty Font 
               Name            =   "�l�r �o����"
               Size            =   9.75
               Charset         =   128
               Weight          =   700
               Underline       =   0   'False
               Italic          =   0   'False
               Strikethrough   =   0   'False
            EndProperty
            Height          =   900
            Left            =   240
            Picture         =   "frmNewListCall.frx":030A
            Style           =   1  '���̨���
            TabIndex        =   56
            Top             =   960
            Width           =   1035
         End
         Begin VB.CommandButton btnExport 
            Caption         =   "�o��"
            BeginProperty Font 
               Name            =   "�l�r �o����"
               Size            =   9.75
               Charset         =   128
               Weight          =   700
               Underline       =   0   'False
               Italic          =   0   'False
               Strikethrough   =   0   'False
            EndProperty
            Height          =   900
            Left            =   1560
            Picture         =   "frmNewListCall.frx":0614
            Style           =   1  '���̨���
            TabIndex        =   57
            Top             =   960
            Width           =   1035
         End
      End
      Begin VB.Frame frmReadCheck 
         BackColor       =   &H00C0E0FF&
         Caption         =   "����������"
         Height          =   3735
         Left            =   240
         TabIndex        =   59
         Top             =   120
         Width           =   14355
         Begin VB.CommandButton btnStaffLocation 
            Caption         =   "��"
            Enabled         =   0   'False
            Height          =   345
            Left            =   13850
            TabIndex        =   53
            Top             =   3240
            Width           =   345
         End
         Begin VB.TextBox txtStaffLocation 
            BackColor       =   &H8000000F&
            Height          =   345
            Left            =   10020
            TabIndex        =   52
            TabStop         =   0   'False
            Top             =   3240
            Width           =   3820
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   16
            Left            =   7320
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   89
            Top             =   3240
            Width           =   2715
            Begin VB.CheckBox chkStaffLocation 
               Caption         =   "��ĒS���ӏ�"
               Height          =   315
               Left            =   120
               TabIndex        =   51
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.TextBox txtCustomerCode 
            BackColor       =   &H8000000F&
            Height          =   345
            Left            =   3060
            MaxLength       =   14
            TabIndex        =   2
            TabStop         =   0   'False
            Top             =   360
            Width           =   4035
         End
         Begin VB.TextBox txtCallUserID 
            BackColor       =   &H8000000F&
            Height          =   345
            Left            =   3060
            TabIndex        =   4
            TabStop         =   0   'False
            Top             =   720
            Width           =   1635
         End
         Begin VB.TextBox txtElectric1Ret 
            BackColor       =   &H8000000F&
            Height          =   345
            Left            =   3060
            TabIndex        =   8
            TabStop         =   0   'False
            Top             =   1080
            Width           =   4035
         End
         Begin VB.CommandButton btnElectric1State 
            Caption         =   "��"
            Enabled         =   0   'False
            Height          =   345
            Left            =   13815
            TabIndex        =   33
            Top             =   1080
            Width           =   345
         End
         Begin VB.CommandButton btnNextCallTimeStart 
            Caption         =   "��"
            Enabled         =   0   'False
            Height          =   345
            Left            =   4560
            TabIndex        =   24
            Top             =   3240
            Width           =   345
         End
         Begin VB.CommandButton btnNextCallTimeEnd 
            Caption         =   "��"
            Enabled         =   0   'False
            Height          =   345
            Left            =   6765
            TabIndex        =   26
            Top             =   3240
            Width           =   345
         End
         Begin VB.TextBox txtNextCallTimeEnd 
            BackColor       =   &H8000000F&
            Height          =   345
            IMEMode         =   3  '�̌Œ�
            Left            =   5265
            TabIndex        =   25
            TabStop         =   0   'False
            Top             =   3240
            Width           =   1635
         End
         Begin VB.TextBox Text13 
            Alignment       =   2  '��������
            BackColor       =   &H8000000F&
            BeginProperty Font 
               Name            =   "�l�r �o����"
               Size            =   12
               Charset         =   128
               Weight          =   700
               Underline       =   0   'False
               Italic          =   0   'False
               Strikethrough   =   0   'False
            EndProperty
            Height          =   360
            IMEMode         =   3  '�̌Œ�
            Left            =   4800
            TabIndex        =   88
            TabStop         =   0   'False
            Text            =   "�`"
            Top             =   3240
            Width           =   555
         End
         Begin VB.TextBox txtNextCallTimeStart 
            BackColor       =   &H8000000F&
            Height          =   345
            IMEMode         =   3  '�̌Œ�
            Left            =   3060
            TabIndex        =   23
            TabStop         =   0   'False
            Top             =   3240
            Width           =   1635
         End
         Begin VB.CommandButton btnReCallDateStart 
            Caption         =   "��"
            Enabled         =   0   'False
            Height          =   345
            Left            =   4560
            TabIndex        =   19
            Top             =   2880
            Width           =   345
         End
         Begin VB.CommandButton btnReCallDateEnd 
            Caption         =   "��"
            Enabled         =   0   'False
            Height          =   345
            Left            =   6765
            TabIndex        =   21
            Top             =   2880
            Width           =   345
         End
         Begin VB.TextBox txtReCallDateEnd 
            BackColor       =   &H8000000F&
            Height          =   345
            IMEMode         =   3  '�̌Œ�
            Left            =   5265
            TabIndex        =   20
            TabStop         =   0   'False
            Top             =   2880
            Width           =   1635
         End
         Begin VB.TextBox Text10 
            Alignment       =   2  '��������
            BackColor       =   &H8000000F&
            BeginProperty Font 
               Name            =   "�l�r �o����"
               Size            =   12
               Charset         =   128
               Weight          =   700
               Underline       =   0   'False
               Italic          =   0   'False
               Strikethrough   =   0   'False
            EndProperty
            Height          =   345
            IMEMode         =   3  '�̌Œ�
            Left            =   4800
            TabIndex        =   87
            TabStop         =   0   'False
            Text            =   "�`"
            Top             =   2880
            Width           =   555
         End
         Begin VB.CommandButton btnSolutionState 
            Caption         =   "��"
            Enabled         =   0   'False
            Height          =   345
            Left            =   13815
            TabIndex        =   42
            Top             =   2160
            Width           =   345
         End
         Begin VB.CommandButton btnGasState 
            Caption         =   "��"
            Enabled         =   0   'False
            Height          =   345
            Left            =   13815
            TabIndex        =   39
            Top             =   1800
            Width           =   345
         End
         Begin VB.CommandButton btnElectric2State 
            Caption         =   "��"
            Enabled         =   0   'False
            Height          =   345
            Left            =   13815
            TabIndex        =   36
            Top             =   1440
            Width           =   345
         End
         Begin VB.TextBox txtSolutionState 
            BackColor       =   &H8000000F&
            Enabled         =   0   'False
            Height          =   345
            IMEMode         =   3  '�̌Œ�
            Left            =   10020
            TabIndex        =   41
            TabStop         =   0   'False
            Top             =   2160
            Width           =   4170
         End
         Begin VB.TextBox txtGasState 
            BackColor       =   &H8000000F&
            Enabled         =   0   'False
            Height          =   345
            IMEMode         =   3  '�̌Œ�
            Left            =   10020
            TabIndex        =   38
            TabStop         =   0   'False
            Top             =   1800
            Width           =   4170
         End
         Begin VB.TextBox txtElectric2State 
            BackColor       =   &H8000000F&
            Enabled         =   0   'False
            Height          =   345
            IMEMode         =   3  '�̌Œ�
            Left            =   10020
            TabIndex        =   35
            TabStop         =   0   'False
            Top             =   1440
            Width           =   4170
         End
         Begin VB.TextBox txtElectric1State 
            BackColor       =   &H8000000F&
            Enabled         =   0   'False
            Height          =   345
            IMEMode         =   3  '�̌Œ�
            Left            =   10020
            TabIndex        =   32
            TabStop         =   0   'False
            Top             =   1080
            Width           =   4170
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   1
            Left            =   360
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   82
            Top             =   360
            Width           =   2715
            Begin VB.CheckBox chkCustomerCode 
               Caption         =   $"frmNewListCall.frx":091E
               Height          =   315
               Left            =   120
               TabIndex        =   1
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   2
            Left            =   360
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   81
            Top             =   720
            Width           =   2715
            Begin VB.CheckBox chkCallUser 
               Caption         =   "�˓d�S����"
               Height          =   315
               Left            =   120
               TabIndex        =   3
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   3
            Left            =   360
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   80
            Top             =   1080
            Width           =   2715
            Begin VB.CheckBox chkElectric1Ret 
               Caption         =   "�˓d�ړI�i�d�C1�j"
               Height          =   315
               Left            =   120
               TabIndex        =   7
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.TextBox txtElectric2Ret 
            BackColor       =   &H8000000F&
            Height          =   345
            Left            =   3060
            TabIndex        =   10
            TabStop         =   0   'False
            Top             =   1440
            Width           =   4035
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   4
            Left            =   360
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   79
            Top             =   1440
            Width           =   2715
            Begin VB.CheckBox chkElectric2Ret 
               Caption         =   "�˓d�ړI�i�d�C2�j"
               Height          =   315
               Left            =   120
               TabIndex        =   9
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.TextBox txtGasRet 
            BackColor       =   &H8000000F&
            Height          =   345
            Left            =   3060
            TabIndex        =   12
            TabStop         =   0   'False
            Top             =   1800
            Width           =   4035
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   5
            Left            =   360
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   78
            Top             =   1800
            Width           =   2715
            Begin VB.CheckBox chkGasRet 
               Caption         =   "�˓d�ړI�i�K�X�j"
               Height          =   315
               Left            =   120
               TabIndex        =   11
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.TextBox txtSolutionRet 
            BackColor       =   &H8000000F&
            Height          =   345
            Left            =   3060
            TabIndex        =   14
            TabStop         =   0   'False
            Top             =   2160
            Width           =   4035
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   6
            Left            =   360
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   77
            Top             =   2160
            Width           =   2715
            Begin VB.CheckBox chkSolutionRet 
               Caption         =   "�˓d�ړI�i�ح���݁j"
               Height          =   315
               Left            =   120
               TabIndex        =   13
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.TextBox txtPhoneNo 
            BackColor       =   &H8000000F&
            Height          =   345
            Left            =   3060
            TabIndex        =   16
            TabStop         =   0   'False
            Top             =   2520
            Width           =   4035
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   7
            Left            =   360
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   76
            Top             =   2520
            Width           =   2715
            Begin VB.CheckBox chkPhoneNo 
               Caption         =   "�d�b�ԍ�"
               Height          =   315
               Left            =   120
               TabIndex        =   15
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   8
            Left            =   360
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   75
            Top             =   2880
            Width           =   2715
            Begin VB.CheckBox chkReCallDate 
               Caption         =   "�ĉ˓d�\���"
               Height          =   315
               Left            =   120
               TabIndex        =   17
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   9
            Left            =   360
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   74
            Top             =   3240
            Width           =   2715
            Begin VB.CheckBox chkNextCallTime 
               Caption         =   "����˓d����"
               Height          =   315
               Left            =   120
               TabIndex        =   22
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.TextBox txtEnterpriseName 
            BackColor       =   &H8000000F&
            Height          =   345
            Left            =   10020
            TabIndex        =   28
            TabStop         =   0   'False
            Top             =   360
            Width           =   4170
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   0
            Left            =   7320
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   73
            Top             =   360
            Width           =   2715
            Begin VB.CheckBox chkEnterpriseName 
               Caption         =   "��Ɩ�"
               Height          =   315
               Left            =   120
               TabIndex        =   27
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.TextBox txtCorporationName 
            BackColor       =   &H8000000F&
            Height          =   345
            Left            =   10020
            TabIndex        =   30
            TabStop         =   0   'False
            Top             =   720
            Width           =   4170
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   10
            Left            =   7320
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   72
            Top             =   720
            Width           =   2715
            Begin VB.CheckBox chkCorporationName 
               Caption         =   "�@�l��"
               Height          =   315
               Left            =   120
               TabIndex        =   29
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   11
            Left            =   7320
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   71
            Top             =   1080
            Width           =   2715
            Begin VB.CheckBox chkElectric1State 
               Caption         =   "�˓d�󋵁i�d�C1�j"
               Height          =   315
               Left            =   120
               TabIndex        =   31
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   12
            Left            =   7320
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   70
            Top             =   1440
            Width           =   2715
            Begin VB.CheckBox chkElectric2State 
               Caption         =   "�˓d�󋵁i�d�C2�j"
               Height          =   315
               Left            =   120
               TabIndex        =   34
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   13
            Left            =   7320
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   69
            Top             =   1800
            Width           =   2715
            Begin VB.CheckBox chkGasState 
               Caption         =   "�˓d�󋵁i�K�X�j"
               Height          =   315
               Left            =   120
               TabIndex        =   37
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   14
            Left            =   7320
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   68
            Top             =   2160
            Width           =   2715
            Begin VB.CheckBox chkSolutionState 
               Caption         =   "�˓d�󋵁i�ح���݁j"
               Height          =   315
               Left            =   120
               TabIndex        =   40
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   15
            Left            =   7320
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   67
            Top             =   2520
            Width           =   2715
            Begin VB.CheckBox chkE1AplForm 
               Caption         =   "�\�����i�d�C1�j"
               Height          =   315
               Left            =   120
               TabIndex        =   43
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   17
            Left            =   7320
            ScaleHeight     =   285
            ScaleWidth      =   2655
            TabIndex        =   66
            Top             =   2880
            Width           =   2715
            Begin VB.CheckBox chkE2AplForm 
               Caption         =   "�\�����i�d�C2�j"
               Height          =   315
               Left            =   120
               TabIndex        =   47
               Top             =   0
               Width           =   5115
            End
         End
         Begin VB.TextBox txtCallUserName 
            BackColor       =   &H8000000F&
            Height          =   345
            Left            =   4680
            TabIndex        =   5
            TabStop         =   0   'False
            Top             =   720
            Width           =   2040
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   18
            Left            =   10035
            ScaleHeight     =   285
            ScaleWidth      =   1200
            TabIndex        =   65
            Top             =   2520
            Width           =   1260
            Begin VB.CheckBox chkE1AplFormWait 
               Caption         =   "��̑�"
               Enabled         =   0   'False
               Height          =   315
               Left            =   120
               TabIndex        =   44
               Top             =   0
               Width           =   5715
            End
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   19
            Left            =   11280
            ScaleHeight     =   285
            ScaleWidth      =   1200
            TabIndex        =   64
            Top             =   2520
            Width           =   1260
            Begin VB.CheckBox chkE1AplFormDone 
               Caption         =   "��̍�"
               Enabled         =   0   'False
               Height          =   315
               Left            =   120
               TabIndex        =   45
               Top             =   0
               Width           =   5235
            End
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   20
            Left            =   12525
            ScaleHeight     =   285
            ScaleWidth      =   1605
            TabIndex        =   63
            Top             =   2520
            Width           =   1660
            Begin VB.CheckBox chkE1AplFormUnnecessary 
               Caption         =   "��̕s�v"
               Enabled         =   0   'False
               Height          =   315
               Left            =   120
               TabIndex        =   46
               Top             =   0
               Width           =   5235
            End
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   21
            Left            =   10035
            ScaleHeight     =   285
            ScaleWidth      =   1200
            TabIndex        =   62
            Top             =   2880
            Width           =   1260
            Begin VB.CheckBox chkE2AplFormWait 
               Caption         =   "��̑�"
               Enabled         =   0   'False
               Height          =   315
               Left            =   120
               TabIndex        =   48
               Top             =   0
               Width           =   5235
            End
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   22
            Left            =   11280
            ScaleHeight     =   285
            ScaleWidth      =   1200
            TabIndex        =   61
            Top             =   2880
            Width           =   1260
            Begin VB.CheckBox chkE2AplFormDone 
               Caption         =   "��̍�"
               Enabled         =   0   'False
               Height          =   315
               Left            =   120
               TabIndex        =   49
               Top             =   0
               Width           =   5235
            End
         End
         Begin VB.PictureBox picDemo 
            Height          =   345
            Index           =   23
            Left            =   12525
            ScaleHeight     =   285
            ScaleWidth      =   1605
            TabIndex        =   60
            Top             =   2880
            Width           =   1660
            Begin VB.CheckBox chkE2AplFormUnnecessary 
               Caption         =   "��̕s�v"
               Enabled         =   0   'False
               Height          =   315
               Left            =   120
               TabIndex        =   50
               Top             =   0
               Width           =   5235
            End
         End
         Begin VB.CommandButton btnUser 
            Enabled         =   0   'False
            Height          =   345
            Left            =   6730
            Picture         =   "frmNewListCall.frx":0930
            Style           =   1  '���̨���
            TabIndex        =   6
            Top             =   721
            Width           =   345
         End
         Begin VB.TextBox txtReCallDateStart 
            BackColor       =   &H8000000F&
            Height          =   345
            Left            =   3060
            TabIndex        =   18
            TabStop         =   0   'False
            Top             =   2880
            Width           =   1635
         End
      End
      Begin MSComDlg.CommonDialog cdlFile 
         Left            =   0
         Top             =   0
         _ExtentX        =   847
         _ExtentY        =   847
         _Version        =   393216
      End
      Begin MSComctlLib.ListView lvwKeep 
         Height          =   2535
         Left            =   240
         TabIndex        =   58
         Top             =   9240
         Width           =   17235
         _ExtentX        =   30401
         _ExtentY        =   4471
         View            =   3
         LabelEdit       =   1
         LabelWrap       =   -1  'True
         HideSelection   =   -1  'True
         FullRowSelect   =   -1  'True
         GridLines       =   -1  'True
         _Version        =   393217
         ForeColor       =   -2147483640
         BackColor       =   -2147483643
         Appearance      =   1
         NumItems        =   16
         BeginProperty ColumnHeader(1) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            Text            =   "ProfileID"
            Object.Width           =   0
         EndProperty
         BeginProperty ColumnHeader(2) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            SubItemIndex    =   1
            Text            =   "�ڋq��"
            Object.Width           =   3351
         EndProperty
         BeginProperty ColumnHeader(3) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            SubItemIndex    =   2
            Text            =   "�d�b�ԍ�"
            Object.Width           =   3351
         EndProperty
         BeginProperty ColumnHeader(4) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            SubItemIndex    =   3
            Object.Width           =   3404
         EndProperty
         BeginProperty ColumnHeader(5) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            SubItemIndex    =   4
            Object.Width           =   3351
         EndProperty
         BeginProperty ColumnHeader(6) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            SubItemIndex    =   5
            Text            =   "�D�挋��"
            Object.Width           =   3351
         EndProperty
         BeginProperty ColumnHeader(7) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            Alignment       =   2
            SubItemIndex    =   6
            Text            =   "�\���(����)"
            Object.Width           =   2822
         EndProperty
         BeginProperty ColumnHeader(8) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            Alignment       =   2
            SubItemIndex    =   7
            Text            =   "�J�n(����)"
            Object.Width           =   2469
         EndProperty
         BeginProperty ColumnHeader(9) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            Alignment       =   2
            SubItemIndex    =   8
            Text            =   "�I��(����)"
            Object.Width           =   2469
         EndProperty
         BeginProperty ColumnHeader(10) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            Alignment       =   2
            SubItemIndex    =   9
            Text            =   "�\���(�x��)"
            Object.Width           =   0
         EndProperty
         BeginProperty ColumnHeader(11) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            Alignment       =   2
            SubItemIndex    =   10
            Text            =   "�J�n(�x��)"
            Object.Width           =   0
         EndProperty
         BeginProperty ColumnHeader(12) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            Alignment       =   2
            SubItemIndex    =   11
            Text            =   "�I��(�x��)"
            Object.Width           =   0
         EndProperty
         BeginProperty ColumnHeader(13) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            SubItemIndex    =   12
            Text            =   "CustomerCode"
            Object.Width           =   0
         EndProperty
         BeginProperty ColumnHeader(14) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            SubItemIndex    =   13
            Text            =   "GroupCode"
            Object.Width           =   0
         EndProperty
         BeginProperty ColumnHeader(15) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            SubItemIndex    =   14
            Text            =   "ListCode"
            Object.Width           =   0
         EndProperty
         BeginProperty ColumnHeader(16) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            SubItemIndex    =   15
            Text            =   "JobCode"
            Object.Width           =   0
         EndProperty
      End
      Begin VB.Label lblErrMsg 
         Alignment       =   1  '�E����
         BackStyle       =   0  '����
         Caption         =   "�G���[�F�������ʌ������������Č��ʂ��\���ł��܂���A�����������i���Ă��������B"
         ForeColor       =   &H000000C0&
         Height          =   255
         Left            =   -4560
         TabIndex        =   86
         Top             =   4080
         Visible         =   0   'False
         Width           =   13275
      End
      Begin VB.Label lblDataCount 
         Alignment       =   1  '�E����
         BackStyle       =   0  '����
         Caption         =   "0���̃f�[�^���Y�����܂���"
         ForeColor       =   &H000000C0&
         Height          =   255
         Left            =   13320
         TabIndex        =   85
         Top             =   4080
         Width           =   4155
      End
   End
End
Attribute VB_Name = "frmNewListCall"
Attribute VB_GlobalNameSpace = False
Attribute VB_Creatable = False
Attribute VB_PredeclaredId = True
Attribute VB_Exposed = False
Option Explicit


'����˓d����sStart
Private Sub btnNextCallTimeEnd_Click()
'    txtNextCallTimeEnd.Text = Format(Now, "yyyy/mm/dd hh:nn:ss")
    txtNextCallTimeEnd.Text = Format(Now, "yyyy/mm/dd")
End Sub


'����˓d����End
Private Sub btnNextCallTimeStart_Click()
'    txtNextCallTimeStart.Text = Format(Now, "yyyy/mm/dd hh:nn:ss")
    txtNextCallTimeStart.Text = Format(Now, "yyyy/mm/dd")
End Sub


'�ĉ˓d�\����J�n
Private Sub btnReCallDateStart_Click()
    txtReCallDateStart.Text = Format(Now, "yyyy/mm/dd")
End Sub


'�ĉ˓d�\����I��
Private Sub btnReCallDateEnd_Click()
    txtReCallDateEnd.Text = Format(Now, "yyyy/mm/dd")
End Sub


'---------------------------------------------------------------------
' �ړI����  -- Keep�ꗗ���X�g - �_�u���N���b�N
' ����      -- �Ȃ�
' �߂�l    -- �Ȃ�
'---------------------------------------------------------------------
Private Sub lvwKeep_DblClick()

On Error GoTo Err_lvwKeep_DblClick
    
    Dim ERRNO As Integer
    Dim SQLERR As Integer
    
    If lvwKeep.ListItems.count = 0 Then Exit Sub
    
#If DEBUG_MODE Then
    TraceLog LOG_DBG, "���X�g�R�[�� Keep�ꗗ �_�u���N���b�N"
#End If
    
    ConnectionOpen gDBString
'    SQL_SP_SELECTCUSTOMERTRN lvwKeep.SelectedItem, gLogonInfo.userID, ERRNO, SQLERR, gCnn
'    If ERRNO <> 0 Then
'        f_MsgPop "���ɑ��̃��[�U���I���ɓ����Ă��܂��B", MSG_ERROR
'        Exit Sub
'    End If
    
    gKeyInfo.CustomerCode = lvwKeep.SelectedItem.SubItems(12)
    gKeyInfo.ReloadFlag = True
    gKeyInfo.ProfileID = lvwKeep.SelectedItem
    '�Ɩ�
    frmMain.txtJobCode = lvwKeep.SelectedItem.SubItems(15)
    '���X�g�R�[�h
    frmMain.txtListCode = lvwKeep.SelectedItem.SubItems(14)
    frmMain.txtGroupCode = lvwKeep.SelectedItem.SubItems(13)
    
    frmMain.txtProfileID = gKeyInfo.ProfileID
    
    '�I������ (TB_Nodial�}��) "0"=���I��, "1"=�I��, "2"=Nodial ("0","1"�͖��g�p�ƂȂ����B�I���R�[�h�̌��̕������Memo�ɑ}�������)
    frmMain.ucPhoneApi1.SetCallSelect gKeyInfo.ProfileID, "2DC"
    
    '�d�bAPI��ProfileID, PhoneNo�ɓ����
    frmMain.ucPhoneApi1.ProfileID = gKeyInfo.ProfileID
    frmMain.ucPhoneApi1.PhoneNo = lvwKeep.SelectedItem.SubItems(2)  '�d�b�ԍ�
    frmMain.ucPhoneApi1.Group = frmMain.txtGroupCode
    
    ISIBFLG = False
    Dim freeDial As String
    'Dialin�ԍ��擾
    freeDial = getDialin(lvwKeep.SelectedItem.SubItems(14), lvwKeep.SelectedItem.SubItems(15))
    If freeDial <> "" Then
        frmMain.ucPhoneApi1.CallingNumber = freeDial
    Else
        frmMain.ucPhoneApi1.CallingNumber = gClientMaster.freeDial
    End If
    
    '�I�𒆂ɂ���
    frmMain.ucPhoneApi1.AppStatus = APSTAT_NEXT
    
    Me.Visible = False
    DoEvents
    
Exit_lvwKeep_DblClick:
    Exit Sub

Err_lvwKeep_DblClick:
    'MsgBox Err.Description
    ErrorBreak "frmListCall.lvwKeep_DblClick"
    Resume Exit_lvwKeep_DblClick

End Sub


'�˓d�\��J�n�I���`�F�b�N
'Private Sub txtReCallDateEnd_LostFocus()
' If DateDiff("d", txtReCallDateStart.Text, txtReCallDateEnd.Text) < 0 Then
'        MsgBox ("�I�����t���J�n���t�ȍ~�̓��t���w�肵�Ă��������B")
'    End If
'End Sub

'����˓d�����`�F�b�N
'Private Sub txtNextCallTimeEnd_LostFocus()
'    If Trim(txtNextCallTimeStart.Text) = "" Or Trim(txtNextCallTimeEnd.Text) = "" Then
'        Exit Sub
'    End If
' If DateDiff("d", txtNextCallTimeStart.Text, txtNextCallTimeEnd.Text) < 0 Then
'        MsgBox ("�I�����t���J�n���t�ȍ~�̓��t���w�肵�Ă��������B")
'    End If
'End Sub


'------------------------------------------------------------------
'�v���V�[�W����:  :btnElectric1State_Click
'����:            :�d�C�P�˓d��
'------------------------------------------------------------------
Private Sub btnElectric1State_Click()
    frmCallState.setTargetCtl 1
    frmCallState.Show 1, Me
End Sub


'------------------------------------------------------------------
'�v���V�[�W����:  :btnElectric2State_Click
'����:            :�d�C�Q�˓d��
'------------------------------------------------------------------
Private Sub btnElectric2State_Click()
    '�d�C1�˓d��
    frmCallState.setTargetCtl 2
    frmCallState.Show 1, Me
End Sub


'------------------------------------------------------------------
'�v���V�[�W����:  :btnGasState_Click
'����:            :�K�X�˓d��
'------------------------------------------------------------------
Private Sub btnGasState_Click()
    frmCallState.setTargetCtl 3
    frmCallState.Show 1, Me
End Sub


'------------------------------------------------------------------
'�v���V�[�W����:  :btnSolutionState_Click
'����:            :�ح���݉˓d��
'------------------------------------------------------------------
Private Sub btnSolutionState_Click()
    frmCallState.setTargetCtl 4
    frmCallState.Show 1, Me
End Sub


'------------------------------------------------------------------
'�v���V�[�W����:  :btnStaffLocation_Click
'����:            :��ĒS���ӏ�
'------------------------------------------------------------------
Private Sub btnStaffLocation_Click()
    frmCallState.setTargetCtl 5
    frmCallState.Show 1, Me
End Sub




'---------------------------------------------------------------------
' �ړI����  -- ���[�U�[�ꗗ
' ����      --
' �߂�l    -- �Ȃ�
'---------------------------------------------------------------------
Private Sub btnUser_Click()

    gSelectUser.UserID = gLogonInfo.UserID
    gSelectUser.UserName = gLogonInfo.UserName
        
    '���[�U�[�ꗗ��\������
    frmUserList.Show (1)
    
    If gQuestionAnswer Then
        Me.txtCallUserID = gSelectUser.UserID
        Me.txtCallUserName = gSelectUser.UserName
    End If

End Sub




'Private dcBrand As Dictionary
'Private dcBrandCode As Dictionary
'Private Const IMPORT_FILE_COMMA_NUM = 10
'
'Private Const IMPORT_RESULT_INS = 1
'Private Const IMPORT_RESULT_UPD = 2
'Private Const IMPORT_RESULT_ERR = 9
'Private Const IMPORT_RESULT_ITEMCOUNT = 1000
'Private Const IMPORT_RESULT_NOBRAND = 2000
'Private Const IMPORT_RESULT_NOCLASSIFIED = 3000
'Private Const IMPORT_ERR_MSG_NOBRAND = "�u�����hCD�G���["
'Private Const IMPORT_ERR_MSG_NOCLASSIFIED = "����CD�G���["
'Private Const IMPORT_ERR_MSG_ITEMCOUNT = "���ڐ��G���["
'Private Const IMPORT_ERR_MSG_SQL = "SQL�G���["
'
'Private Const LVM_FIRST = &H1000&
'Private Const LVM_SETCOLUMNWIDTH = (LVM_FIRST + 30)
'Private Const LVSCW_AUTOSIZE_USEHEADER = -2


'Private Sub btnClassification_Click()
'
'    frmProductClassification.Show (1)
'
'End Sub

'------------------------------------------------------------------
'�v���V�[�W����:  :btnExport_Click
'����:            :�˓d���X�g�\�����e�̏o��
'------------------------------------------------------------------
Private Sub btnExport_Click()
On Error GoTo Err_btnExport_Click

    Dim i, j                As Long
    Dim strFileTmp          As String
    Dim strPath             As String
    Dim FileNumber          As Integer
    Dim strLine             As String
    Dim strHeader           As String
    
    '���`�F�b�N
    If lvwCallList.ItemCount = 0 Then
        f_MsgPop "�o�͑Ώۂ�����܂���B", MSG_ERROR
        Exit Sub
    End If
    
    cdlFile.CancelError = True
    cdlFile.Filter = "CSV�t�@�C��(*.csv)|*.csv|"
    cdlFile.FilterIndex = 1
    cdlFile.Flags = cdlOFNOverwritePrompt
    cdlFile.fileName = f_GetDeskTopFolder & "\" & "�˓d���X�g�ꗗ_" & Format(Now, "yyyymmddhhnnss") & ".csv" & ""
    cdlFile.ShowSave
    strPath = cdlFile.fileName
    
    DoEvents
    
    '�o�̓w�b�_�[����ǂݍ���
    strHeader = LoadSQL(App.path & "\" & NEW_EXPORT_HEADER)
    strHeader = Replace(strHeader, vbCrLf, "")
    FileNumber = FreeFile
    Open strPath For Output As #FileNumber
    '�w�b�_�[�̏����o��
    Print #FileNumber, strHeader
    
    Dim str As String
    Dim vntData     As Variant
    
    For i = 0 To lvwCallList.ItemCount - 1
        str = ""
        str = lvwCallList.ItemGet(CLng(i))
        vntData = Split(str, "" & vbTab)
        
        'strLine = lvwProduct.ListItems(i).SubItems(1)
        strLine = """" & vntData(1) & """"
        strLine = strLine & "," & """" & vntData(2) & """"
        strLine = strLine & "," & """" & vntData(3) & """"
        strLine = strLine & "," & """" & vntData(4) & """"
        strLine = strLine & "," & """" & vntData(5) & """"
        strLine = strLine & "," & """" & vntData(6) & """"
        strLine = strLine & "," & """" & vntData(7) & """"
        strLine = strLine & "," & """" & vntData(8) & """"
        strLine = strLine & "," & """" & vntData(9) & """"
        strLine = strLine & "," & """" & vntData(10) & """"
        strLine = strLine & "," & """" & vntData(11) & """"
        strLine = strLine & "," & """" & vntData(12) & """"
        strLine = strLine & "," & """" & vntData(13) & """"
        strLine = strLine & "," & """" & vntData(14) & """"
        strLine = strLine & "," & """" & vntData(15) & """"
        strLine = strLine & "," & """" & vntData(16) & """"
        strLine = strLine & "," & """" & vntData(17) & """"
        strLine = strLine & "," & """" & vntData(18) & """"
        strLine = strLine & "," & """" & vntData(19) & """"
        strLine = strLine & "," & """" & vntData(20) & """"
        strLine = strLine & "," & """" & vntData(21) & """"
        strLine = strLine & "," & """" & vntData(22) & """"
        strLine = strLine & "," & """" & vntData(23) & """"
        strLine = strLine & "," & """" & vntData(24) & """"
        strLine = strLine & "," & """" & vntData(25) & """"
        strLine = strLine & "," & """" & vntData(26) & """"
        strLine = strLine & "," & """" & vntData(27) & """"
        strLine = strLine & "," & """" & vntData(28) & """"
        strLine = strLine & "," & """" & vntData(29) & """"
        strLine = strLine & "," & """" & vntData(30) & """"
        strLine = strLine & "," & """" & vntData(31) & """"
        strLine = strLine & "," & """" & vntData(32) & """"
        strLine = strLine & "," & """" & vntData(33) & """"
        strLine = strLine & "," & """" & vntData(34) & """"
        strLine = strLine & "," & """" & vntData(35) & """"
        strLine = strLine & "," & """" & vntData(36) & """"
        strLine = strLine & "," & """" & vntData(37) & """"
        strLine = strLine & "," & """" & vntData(38) & """"
        strLine = strLine & "," & """" & vntData(39) & """"
        strLine = strLine & "," & """" & vntData(40) & """"
        strLine = strLine & "," & """" & vntData(41) & """"
        strLine = strLine & "," & """" & vntData(42) & """"
        strLine = strLine & "," & """" & vntData(43) & """"
        strLine = strLine & "," & """" & vntData(44) & """"
        strLine = strLine & "," & """" & vntData(45) & """"
        strLine = strLine & "," & """" & vntData(46) & """"
        strLine = strLine & "," & """" & vntData(47) & """"

        Print #FileNumber, strLine

    Next i
    
    Close #FileNumber
    
    f_MsgPop "�o�͂��������܂����B", MSG_COMPLETE

Exit_btnExport_Click:
    Exit Sub
Err_btnExport_Click:
    ErrorBreak "btnExport_Click"
    Resume Exit_btnExport_Click
End Sub


'------------------------------------------------------------------
'�v���V�[�W����:  :btnSearch_Click
'����:            :�˓d���X�g�������
'------------------------------------------------------------------
Private Sub btnSearch_Click()
On Error GoTo Err_btnSearch_Click

    Dim cnn As New ADODB.Connection
    Dim rs As New ADODB.Recordset
    Dim SQL, countSQL As String
    Dim SQL_WHERE As String: SQL_WHERE = ""
    Dim i, j As Integer
    Dim ItemX As ListItem
    Dim strConditionMsg As String
    Dim rowCount As String: rowCount = 0
    
    '�������A������
    OnceFlg = False


    lvwCallList.Clear
    '�����pDIC������
    ROW_COMPLETE_DIC.RemoveAll
    '�e�L�^DIC������
    ROW_PARENT_DIC.RemoveAll
    lblDataCount.Caption = "0���̃f�[�^���Y�����܂���"
    lblErrMsg.Visible = False
    
    lvwCallList.CNS = gDBString
    
'
'    cnn.Open gDBString
'
    SQL_WHERE = mMakeWhereSQL
     

    '���sSQL�t�@�C�������[�h
    SQL = LoadSQL(App.path & "\" & SQL_GET_CALL_FIELD)
    '��������
    lvwCallList.FIELDNAME = SQL
    
    If SQL_WHERE = "" Then
'        SQL = Replace(SQL, RE_FIND_WHERE_HOLDER, "")
'        lvwCallList.WHERE = ""
        lvwCallList.WHERE = "ClientCode = " & gLogonInfo.ClientCode
    Else
'        SQL_WHERE = " WHERE " & SQL_WHERE
'        SQL = Replace(SQL, RE_FIND_WHERE_HOLDER, SQL_WHERE) & " "
        '��������
        lvwCallList.WHERE = "ClientCode = " & gLogonInfo.ClientCode & " AND " & SQL_WHERE
    End If
    
    '�\�[�g�L�[�w��
    lvwCallList.ORDERBY = "order by data043,data046,data001"
    '�e�[�u���w��
    lvwCallList.TABLENAME = "TX_KEPOCLIENTCALLINFO"
    
    lvwCallList.SetImageList App.path & "\newlistimg.bmp"
    
    
    '�����`�F�b�N
    countSQL = "select count(1) as rCount from TX_KEPOCLIENTCALLINFO WHERE " & lvwCallList.WHERE
    cnn.Open gDBString
    rs.Open countSQL, cnn
    Do Until rs.EOF
      rowCount = rs.Fields("rCount").Value & ""
      rs.MoveNext
    Loop
    'lblErrMsg �����`�F�b�N
    lblDataCount.Caption = "0���̃f�[�^���Y�����܂���"
    lblErrMsg.Visible = False
    If CLng(rowCount) > MAX_ROW_NUM Then
       lblErrMsg.Visible = True
       lblDataCount.Caption = Replace(lblDataCount.Caption, "0", CommaSplitFmt(rowCount))
       Exit Sub
    Else
        lblErrMsg.Visible = False
        lblDataCount.Caption = Replace(lblDataCount.Caption, "0", CommaSplitFmt(rowCount))
        '�������{
        lvwCallList.Search
    End If
    
Exit_btnSearch_Click:
    Exit Sub
Err_btnSearch_Click:
    MsgBox Err.Description
    Resume Exit_btnSearch_Click
End Sub

'���q�l�ԍ��L��
Private Sub chkCustomerCode_Click()
    If chkCustomerCode.Value = 0 Then
        SetObjectEnable txtCustomerCode, False
    Else
        SetObjectEnable txtCustomerCode, True
    End If
End Sub

'��Ɩ�
Private Sub chkEnterpriseName_Click()
    If chkEnterpriseName.Value = 0 Then
        SetObjectEnable txtEnterpriseName, False
    Else
        SetObjectEnable txtEnterpriseName, True
    End If
End Sub

'�˓d�S����
Private Sub chkCallUser_Click()
    If chkCallUser.Value = 0 Then
        SetObjectEnable txtCallUserID, False
        SetObjectEnable txtCallUserName, False
        btnUser.Enabled = False
    Else
        SetObjectEnable txtCallUserID, True
        SetObjectEnable txtCallUserName, True
        btnUser.Enabled = True
    End If
End Sub

'�@�l��
Private Sub chkCorporationName_Click()
    If chkCorporationName.Value = 0 Then
        SetObjectEnable txtCorporationName, False
    Else
        SetObjectEnable txtCorporationName, True
    End If
End Sub

'�d�C1-�˓d�ړI
Private Sub chkElectric1Ret_Click()
    If chkElectric1Ret.Value = 0 Then
        SetObjectEnable txtElectric1Ret, False
    Else
        SetObjectEnable txtElectric1Ret, True
    End If
End Sub

'�d�C2-�˓d�ړI
Private Sub chkElectric2Ret_Click()
    If chkElectric2Ret.Value = 0 Then
        SetObjectEnable txtElectric2Ret, False
    Else
        SetObjectEnable txtElectric2Ret, True
    End If
End Sub

'�K�X-�˓d�ړI
Private Sub chkGasRet_Click()
    If chkGasRet.Value = 0 Then
        SetObjectEnable txtGasRet, False
    Else
        SetObjectEnable txtGasRet, True
    End If
End Sub



'�ح����-�˓d�ړI
Private Sub chkSolutionRet_Click()
    If chkSolutionRet.Value = 0 Then
        SetObjectEnable txtSolutionRet, False
    Else
        SetObjectEnable txtSolutionRet, True
    End If
End Sub


'�d�C1�|�˓d��
Private Sub chkElectric1State_Click()
    If chkElectric1State.Value = 0 Then
        SetObjectEnable txtElectric1State, False
        SetObjectEnable btnElectric1State, False
    Else
        SetObjectEnable txtElectric1State, True
        SetObjectEnable btnElectric1State, True
    End If
End Sub

'�d�C2�|�˓d��
Private Sub chkElectric2State_Click()
    If chkElectric2State.Value = 0 Then
        SetObjectEnable txtElectric2State, False
        SetObjectEnable btnElectric2State, False
    Else
        SetObjectEnable txtElectric2State, True
        SetObjectEnable btnElectric2State, True
    End If
End Sub

'�K�X�|�˓d��
Private Sub chkGasState_Click()
    If chkGasState.Value = 0 Then
        SetObjectEnable txtGasState, False
        SetObjectEnable btnGasState, False
    Else
        SetObjectEnable txtGasState, True
        SetObjectEnable btnGasState, True
    End If
End Sub

'�\�����[�V�����|�˓d��
Private Sub chkSolutionState_Click()
    If chkSolutionState.Value = 0 Then
        SetObjectEnable txtSolutionState, False
        SetObjectEnable btnSolutionState, False
    Else
        SetObjectEnable txtSolutionState, True
        SetObjectEnable btnSolutionState, True
    End If
End Sub

'�d�b�ԍ�
Private Sub chkPhoneNo_Click()
    If chkPhoneNo.Value = 0 Then
        SetObjectEnable txtPhoneNo, False
        
    Else
        SetObjectEnable txtPhoneNo, True
    End If
End Sub


'�ĉ˓d�\���
Private Sub chkReCallDate_Click()
    If chkReCallDate.Value = 0 Then
        SetObjectEnable txtReCallDateStart, False
        SetObjectEnable txtReCallDateEnd, False
        SetObjectEnable btnReCallDateStart, False
        SetObjectEnable btnReCallDateEnd, False
        
    Else
        SetObjectEnable txtReCallDateStart, True
        SetObjectEnable txtReCallDateEnd, True
        SetObjectEnable btnReCallDateStart, True
        SetObjectEnable btnReCallDateEnd, True
    End If
End Sub

'����˓d����
Private Sub chkNextCallTime_Click()
    If chkNextCallTime.Value = 0 Then
        SetObjectEnable txtNextCallTimeStart, False
        SetObjectEnable txtNextCallTimeEnd, False
        SetObjectEnable btnNextCallTimeStart, False
        SetObjectEnable btnNextCallTimeEnd, False
        
    Else
        SetObjectEnable txtNextCallTimeStart, True
        SetObjectEnable txtNextCallTimeEnd, True
        SetObjectEnable btnNextCallTimeStart, True
        SetObjectEnable btnNextCallTimeEnd, True
    End If
End Sub

'�\�����d�C�P
Private Sub chkE1AplForm_Click()
    If chkE1AplForm.Value = 0 Then
        chkE1AplFormWait.Enabled = False
        chkE1AplFormDone.Enabled = False
        chkE1AplFormUnnecessary.Enabled = False
    Else
        chkE1AplFormWait.Enabled = True
        chkE1AplFormDone.Enabled = True
        chkE1AplFormUnnecessary.Enabled = True
    End If
End Sub

'�\�����d�C2
Private Sub chkE2AplForm_Click()
    If chkE2AplForm.Value = 0 Then
        chkE2AplFormWait.Enabled = False
        chkE2AplFormDone.Enabled = False
        chkE2AplFormUnnecessary.Enabled = False
    Else
        chkE2AplFormWait.Enabled = True
        chkE2AplFormDone.Enabled = True
        chkE2AplFormUnnecessary.Enabled = True
    End If
End Sub

'��ĒS���ӏ�
Private Sub chkStaffLocation_Click()
    If chkStaffLocation.Value = 0 Then
        SetObjectEnable txtStaffLocation, False
        btnStaffLocation.Enabled = False
    Else
        SetObjectEnable txtStaffLocation, True
        btnStaffLocation.Enabled = True
    End If
End Sub


'------------------------------------------------------------------
'�v���V�[�W����:  :mMakeWhereSQL
'����:            :���������쐬
'�߂�l           :Variant�^
'------------------------------------------------------------------
Private Function mMakeWhereSQL() As String
    Dim WhereSQL As String: WhereSQL = ""
    Dim RunMod As String: RunMod = " AND "
    Dim i As Integer
    '���s���[�h
    If optMod.item(0).Value = True Then
        RunMod = " OR "
    End If
    
    
    '���q�l�ԍ��I�������ꍇ
    If chkCustomerCode.Value <> 0 Then
        WhereSQL = WhereSQL & " Data001 Like '%" & txtCustomerCode.Text & "%' " & RunMod
    End If
    
    '��Ɩ�  ��Ɩ��i�����j�F�J���f��������Ё@���@�J���f���Ō�������
    '�c�a���̒n�_���A Data041
    '���v�ꏊ�Z���    Data096
    '�x����Z���      Data097
    '�x���Җ��A       Data082
    If chkEnterpriseName.Value <> 0 Then
        WhereSQL = WhereSQL & "(Data041 Like '%" & txtEnterpriseName.Text & "%' OR Data096 Like '%" & txtEnterpriseName.Text & _
        "%' OR Data097 Like '%" & txtEnterpriseName.Text & "%' OR Data082 Like '%" & txtEnterpriseName.Text & "%')" & RunMod
    End If
    
    '�I�������˓d�҂ɍ��v����s�S��  Data098
    If chkCallUser.Value <> 0 Then
        WhereSQL = WhereSQL & "(Data098 = '" & txtCallUserName.Text & "' OR Data098 = '" & txtCallUserID.Text & "')" & RunMod
    End If
    
    '�@�l�� Data043
    If chkCorporationName.Value <> 0 Then
        WhereSQL = WhereSQL & " Data043 Like '%" & txtCorporationName.Text & "%' " & RunMod
    End If
    
    '�d�C1-�˓d�ړI chkElectric1Ret_Click �������镶�����܂ލs���ׂ�
    If chkElectric1Ret.Value <> 0 Then
        WhereSQL = WhereSQL & " Data049 Like '%" & txtElectric1Ret.Text & "%' " & RunMod
    End If
    
    '�d�C-�˓d�ړI chkElectric1Ret_Click �������镶�����܂ލs���ׂ�
    If chkElectric2Ret.Value <> 0 Then
        WhereSQL = WhereSQL & " Data050 Like '%" & txtElectric2Ret.Text & "%' " & RunMod
    End If
    
    
    '�K�X-�˓d�ړI  chkGasRet �������镶�����܂ލs���ׂ�
    If chkGasRet.Value <> 0 Then
        WhereSQL = WhereSQL & " Data051 Like '%" & txtGasRet.Text & "%' " & RunMod
    End If
        
    '�ح����-�˓d�ړI chkElectric1Ret_Click �������镶�����܂ލs���ׂ�
    If chkSolutionRet.Value <> 0 Then
        WhereSQL = WhereSQL & " Data052 Like '%" & txtSolutionRet.Text & "%' " & RunMod
    End If
    
    '�d�C1�˓d��
    If chkElectric1State <> 0 Then
        Dim tmpSplit() As String
        Dim j As Integer
        Dim orWhere As String: orWhere = ""
        tmpSplit = Split(txtElectric1State.Text, ",")
        If UBound(tmpSplit) >= 0 Then
            For j = 0 To UBound(tmpSplit)
                If j <> UBound(tmpSplit) Then
                    orWhere = orWhere & "'" & Convert2CallCode(tmpSplit(j)) & "',"
                Else
                    orWhere = orWhere & "'" & Convert2CallCode(tmpSplit(j)) & "'"
                End If
            Next
        End If
        If orWhere = "" Then orWhere = "''"
        WhereSQL = WhereSQL & " Den1CallState IN (" & orWhere & ") " & RunMod
    End If
    
    '�d�C2�˓d��
    If chkElectric2State <> 0 Then
        Dim tmpSplit2() As String
        Dim h As Integer
        Dim orWhere2 As String: orWhere2 = ""
        tmpSplit2 = Split(txtElectric2State.Text, ",")
        If UBound(tmpSplit2) >= 0 Then
            For h = 0 To UBound(tmpSplit2)
                If h <> UBound(tmpSplit2) Then
                    orWhere2 = orWhere2 & "'" & Convert2CallCode(tmpSplit2(h)) & "',"
                Else
                    orWhere2 = orWhere2 & "'" & Convert2CallCode(tmpSplit2(h)) & "'"
                End If
            Next
        End If
        If orWhere2 = "" Then orWhere2 = "''"
        WhereSQL = WhereSQL & " Den2CallState IN (" & orWhere2 & ") " & RunMod
    End If
    
    
    '�K�X�˓d��
    If chkGasState <> 0 Then
        Dim tmpSplit3() As String
        Dim k As Integer
        Dim orWhere3 As String: orWhere3 = ""
        tmpSplit3 = Split(txtGasState.Text, ",")
        If UBound(tmpSplit3) >= 0 Then
            For k = 0 To UBound(tmpSplit3)
                If k <> UBound(tmpSplit3) Then
                    orWhere3 = orWhere3 & "'" & Convert2CallCode(tmpSplit3(k)) & "',"
                Else
                    orWhere3 = orWhere3 & "'" & Convert2CallCode(tmpSplit3(k)) & "'"
                End If
            Next
        End If
        If orWhere3 = "" Then orWhere3 = "''"
        WhereSQL = WhereSQL & " GasCallState IN (" & orWhere3 & ") " & RunMod
    End If
    
    '�ح���݉˓d��
    If chkSolutionState <> 0 Then
        Dim tmpSplit4() As String
        Dim f As Integer
        Dim orWhere4 As String: orWhere4 = ""
        tmpSplit4 = Split(txtSolutionState.Text, ",")
        If UBound(tmpSplit4) >= 0 Then
            For f = 0 To UBound(tmpSplit4)
                If f <> UBound(tmpSplit4) Then
                    orWhere4 = orWhere4 & "'" & Convert2CallCode(tmpSplit4(f)) & "',"
                Else
                    orWhere4 = orWhere4 & "'" & Convert2CallCode(tmpSplit4(f)) & "'"
                End If
            Next
        End If
        If orWhere4 = "" Then orWhere4 = "''"
        WhereSQL = WhereSQL & " SoluCallState IN (" & orWhere4 & ") " & RunMod
    End If
    
    '�d�b�ԍ���\�� �����̓d�b�ԍ��ƍ��v����s�S��
    If chkPhoneNo.Value <> 0 Then
        WhereSQL = WhereSQL & " Data002 ='" & txtPhoneNo.Text & "' " & RunMod
    End If
    
    
    '�ĉ˓d���� �N�������͔N�������܂ލsor�ȑO�̍sor�ȍ~�̍s�������ł���悤��
    If chkReCallDate.Value <> 0 Then
        'WhereSQL = WhereSQL & " (ReCallDate >='" & txtReCallDateStart.Text & "' And " & " ReCallDate <='" & txtReCallDateEnd.Text & "') " & RunMod
        '�ĉ˓d�J�n��
        Dim callstart As String: callstart = IIf(Trim(txtReCallDateStart.Text) = "", "2000/01/01", txtReCallDateStart.Text)
        '�ĉ˓d�I����
        Dim callend As String: callend = IIf(Trim(txtReCallDateEnd.Text) = "", "2300/12/31", txtReCallDateEnd.Text)
        WhereSQL = WhereSQL & " (ReCallDate >='" & callstart & "' And " & " ReCallDate <='" & callend & "') " & RunMod
    End If
    
    '����˓d����(�d�C�P�A�d�C�Q�A�K�X�A�\�����[�V����)
    If chkNextCallTime.Value <> 0 Then
        'Default�J�n
        Dim defaultStart As String: defaultStart = IIf(Trim(txtNextCallTimeStart.Text) = "", "2000/01/01", txtNextCallTimeStart.Text)
        'Default�I��
        Dim defaultEnd As String: defaultEnd = IIf(Trim(txtNextCallTimeEnd.Text) = "", "2300/12/31", txtNextCallTimeEnd.Text)
        

'       WhereSQL = WhereSQL & " ((Den1NextCallDate >='" & txtNextCallTimeStart.Text & "' And " & " Den1NextCallDate <='" & txtNextCallTimeEnd.Text & "') OR " & _
'                              " (Den2NextCallDate >='" & txtNextCallTimeStart.Text & "' And " & " Den2NextCallDate <='" & txtNextCallTimeEnd.Text & "') OR " & _
'                              " (GasNextCallDate >='" & txtNextCallTimeStart.Text & "' And " & " GasNextCallDate <='" & txtNextCallTimeEnd.Text & "') OR " & _
'                              " (SoluNextCallDate >='" & txtNextCallTimeStart.Text & "' And " & " SoluNextCallDate <='" & txtNextCallTimeEnd.Text & "'))" & _
'                              "" & RunMod
    
        WhereSQL = WhereSQL & " ((Den1NextCallDate >='" & defaultStart & "' And " & " Den1NextCallDate <='" & defaultEnd & "') OR " & _
                              " (Den2NextCallDate >='" & defaultStart & "' And " & " Den2NextCallDate <='" & defaultEnd & "') OR " & _
                              " (GasNextCallDate >='" & defaultStart & "' And " & " GasNextCallDate <='" & defaultEnd & "') OR " & _
                              " (SoluNextCallDate >='" & defaultStart & "' And " & " SoluNextCallDate <='" & defaultEnd & "'))" & _
                              "" & RunMod
    End If
    
    '�d�C1�\����
    If chkE1AplForm.Value <> 0 Then
        '��̑҂�
        If chkE1AplFormWait.Value <> 0 Then
            WhereSQL = WhereSQL & " (Den1CallApplyLetter ='" & ConvertAPL2Code(chkE1AplFormWait.Caption) & "') " & "OR "
        End If
        
        If chkE1AplFormDone.Value <> 0 Then
            WhereSQL = WhereSQL & " (Den1CallApplyLetter ='" & ConvertAPL2Code(chkE1AplFormDone.Caption) & "') " & "OR "
        End If
        
        If chkE1AplFormUnnecessary.Value <> 0 Then
            WhereSQL = WhereSQL & " (Den1CallApplyLetter ='" & ConvertAPL2Code(chkE1AplFormUnnecessary.Caption) & "') " & "OR "
        End If
        
        '�Ō�OR������
        If Right(WhereSQL, Len("OR ")) = "OR " Then
            WhereSQL = Left(WhereSQL, Len(WhereSQL) - Len("OR ")) & RunMod
        Else
            WhereSQL = WhereSQL
        End If
    End If
    
    '�d�C2�\����
    If chkE2AplForm.Value <> 0 Then
        '��̑҂�
        If chkE2AplFormWait.Value <> 0 Then
            WhereSQL = WhereSQL & " (Den2CallApplyLetter ='" & ConvertAPL2Code(chkE2AplFormWait.Caption) & "') " & "OR "
        End If
        
        If chkE2AplFormDone.Value <> 0 Then
            WhereSQL = WhereSQL & " (Den2CallApplyLetter ='" & ConvertAPL2Code(chkE2AplFormDone.Caption) & "') " & "OR "
        End If
        
        If chkE2AplFormUnnecessary.Value <> 0 Then
            WhereSQL = WhereSQL & " (Den2CallApplyLetter ='" & ConvertAPL2Code(chkE2AplFormUnnecessary.Caption) & "') " & "OR "
        End If
        
        '�Ō�OR������
        If Right(WhereSQL, Len("OR ")) = "OR " Then
            WhereSQL = Left(WhereSQL, Len(WhereSQL) - Len("OR ")) & RunMod
        Else
            WhereSQL = WhereSQL
        End If
    End If
    
    '��ĒS���ӏ�
    If chkStaffLocation.Value <> 0 Then
        'whereSQL = whereSQL & " Data009 Like '%" & txtStaffLocation.Text & "%' " & RunMod
        If chkStaffLocation.Value <> 0 Then
            Dim tmpSplit5() As String
            Dim g As Integer
            Dim orWhere5 As String: orWhere5 = ""
            tmpSplit5 = Split(txtStaffLocation.Text, ",")
            If UBound(tmpSplit5) >= 0 Then
                For g = 0 To UBound(tmpSplit5)
                    If g <> UBound(tmpSplit5) Then
'                        orWhere5 = orWhere5 & "'" & Convert2CallCode(tmpSplit5(g)) & "',"
                        orWhere5 = orWhere5 & "'" & convertMethod.mConvertValue2Code(tmpSplit5(g)) & "',"
                    Else
                        'orWhere5 = orWhere5 & "'" & Convert2CallCode(tmpSplit5(g)) & "'"
                        orWhere5 = orWhere5 & "'" & convertMethod.mConvertValue2Code(tmpSplit5(g)) & "'"
                    End If
                Next
            End If
            If orWhere5 = "" Then orWhere5 = "''"
            WhereSQL = WhereSQL & " Data009 IN (" & orWhere5 & ") " & RunMod
        End If
    End If

    If WhereSQL <> "" Then
       i = InStrRev(WhereSQL, RunMod)
       WhereSQL = Left(WhereSQL, i)
    End If
    mMakeWhereSQL = WhereSQL
End Function


Private Sub Form_Load()
    'DB�A�N�Z�X���擾
    iniGetValueSQLServer gDBString
    
    '�R�[�h�}�X�^�擾
    Load_CodeMaster
    
    '�R�[�h�}�X�^�[���[�h-�d�C1�\����
    Load_APLForm APL_CODE_1, FC_CODE_1, CATEGORY_1
    
    '��ĒS���ӏ�CSV
    ReadGroupCode App.path & "\" & GroupCodeCsv
    
    '���X�g�r���[������
    makeCallListHeader

    '���X�g�r���[�w�b�_�F�ݒ�
    'SubClass lvwCallList.hwnd, AddressOf LVWndProc
    '���X�g�r���[�s�w�i�ύX
    'SubClassForRowColor Me.Picture1.hWnd
End Sub

'Private Sub Form_QueryUnload(Cancel As Integer, UnloadMode As Integer)
'    UnSubClassForRowColor Me.Picture1.hWnd
'End Sub


'�R�[�����X�g�r���[�̃w�b�_�[�ݒ�
'------------------------------------------------------------------
'�v���V�[�W����:  :makeCallListHeader
'����:            :�˓d���X�g�r���[�w�b�_�[�ݒ�
'------------------------------------------------------------------
Private Sub makeCallListHeader()
    
    lvwCallList.Clear
    lvwCallList.ColumnAdd 0, 0, 0, ""
    lvwCallList.ColumnAdd 1, 0, 0, ""
    'TB_CLIENTDATA
    lvwCallList.ColumnAdd 1, 0, 135, "���q���ܔԍ�"
    'TB_CLIENTDATA
    lvwCallList.ColumnAdd 2, 0, 165, "�n�_��"
    'TB_CLIENTDATA
    lvwCallList.ColumnAdd 3, 0, 165, "�n�_���J�i"
    'TB_CLIENTDATA
    lvwCallList.ColumnAdd 4, 0, 165, "�g�p���"
    'TB_CLIENTDATA
    lvwCallList.ColumnAdd 5, 0, 165, "�˓d�Ώے������R"
    
    lvwCallList.ColumnAdd 6, 0, 165, "��ĒS���ӏ�"
    
    'TX_KEPCallInfo
    lvwCallList.ColumnAdd 7, 0, 165, "�˓d�ړI-(�d�C�P)"
    '��
    lvwCallList.SetColumnColor 7, vbRed
    'lvwCallList.SetColumnColor 7, RGB(255, 0, 0)
    
    'TX_KEPCallInfo
    lvwCallList.ColumnAdd 8, 0, 165, "�˓d�ړI-(�d�C�Q)"
    '�I�����W
    lvwCallList.SetColumnColor 8, RGB(255, 165, 0)
    
    'TX_KEPCallInfo
    lvwCallList.ColumnAdd 9, 0, 165, "�˓d�ړI-(�K�X)"
    '��
    lvwCallList.SetColumnColor 9, vbBlue
    
    'TX_KEPCallInfo
    lvwCallList.ColumnAdd 10, 0, 165, "�˓d�ړI-(�ح����)"
    '��
    lvwCallList.SetColumnColor 10, RGB(91, 162, 97)
    
    'TX_KEPCallInfo
    lvwCallList.ColumnAdd 11, 0, 165, "�Ǝ�"
    'TX_KEPCallInfo
    lvwCallList.ColumnAdd 12, 0, 165, "�e�q����"
    
    'TB_ClientDaata
    lvwCallList.ColumnAdd 13, 0, 165, "�@�l��"
    'TB_ClientDaata
    lvwCallList.ColumnAdd 14, 0, 165, "�˓d�S����"
    'TB_ClientDaata
    lvwCallList.ColumnAdd 15, 0, 165, "��\TEL"
    'TB_ClientDaata
    lvwCallList.ColumnAdd 16, 0, 165, "�x����TEL"
    'TB_ClientDaata
    lvwCallList.ColumnAdd 17, 0, 165, "���v�ꏊ�X�֔ԍ�"
    'TB_ClientDaata
    lvwCallList.ColumnAdd 18, 0, 165, "���v�ꏊ�Z��"
    'TB_ClientDaata
    lvwCallList.ColumnAdd 19, 0, 165, "�x����X�֔ԍ�"
    'TB_ClientDaata
    lvwCallList.ColumnAdd 20, 0, 165, "�x����Z��"
    'TB_ClientDaata
    lvwCallList.ColumnAdd 21, 0, 165, "�x���Җ�"
    lvwCallList.ColumnAdd 22, 0, 165, "�ĉ˓d��"
    lvwCallList.ColumnAdd 23, 0, 165, "�˓d���ԑсi�n�j"
    lvwCallList.ColumnAdd 24, 0, 165, "�˓d���ԑсi�I�j"
    
    'TX_ACTIVEHIS
    lvwCallList.ColumnAdd 25, 0, 165, "�˓d����-(�d�C�P)"
    '��
    lvwCallList.SetColumnColor 25, vbRed
    
    lvwCallList.ColumnAdd 26, 0, 165, "�˓d��-(�d�C�P)"
    '��
    lvwCallList.SetColumnColor 26, vbRed
    
    lvwCallList.ColumnAdd 27, 0, 165, "�˓d����-(�d�C�P)"
    '��
    lvwCallList.SetColumnColor 27, vbRed
    
    lvwCallList.ColumnAdd 28, 0, 210, "����˓d����-(�d�C�P)"
    '��
    lvwCallList.SetColumnColor 28, vbRed
    
    lvwCallList.ColumnAdd 29, 0, 165, "�s�����R-(�d�C�P)"
    '��
    lvwCallList.SetColumnColor 29, vbRed
    
    lvwCallList.ColumnAdd 30, 0, 210, "�\������̓�-(�d�C�P)"
    '��
    lvwCallList.SetColumnColor 30, vbRed
    
    
    lvwCallList.ColumnAdd 31, 0, 165, "�˓d����-(�d�C�Q)"
    '�I�����W
    lvwCallList.SetColumnColor 31, RGB(255, 165, 0)
    
    lvwCallList.ColumnAdd 32, 0, 165, "�˓d��-(�d�C�Q)"
    '�I�����W
    lvwCallList.SetColumnColor 32, RGB(255, 165, 0)
    
    lvwCallList.ColumnAdd 33, 0, 165, "�˓d����-(�d�C�Q)"
    '�I�����W
    lvwCallList.SetColumnColor 33, RGB(255, 165, 0)
    
    lvwCallList.ColumnAdd 34, 0, 210, "����˓d����-(�d�C�Q)"
    '�I�����W
    lvwCallList.SetColumnColor 34, RGB(255, 165, 0)
    
    lvwCallList.ColumnAdd 35, 0, 165, "�s�����R-(�d�C�Q)"
    '�I�����W
    lvwCallList.SetColumnColor 35, RGB(255, 165, 0)
    
    lvwCallList.ColumnAdd 36, 0, 210, "�\������̓�-(�d�C�Q)"
    '�I�����W
    lvwCallList.SetColumnColor 36, RGB(255, 165, 0)
    
    lvwCallList.ColumnAdd 37, 0, 165, "�˓d����-(�K�X)"
    '��
    lvwCallList.SetColumnColor 37, vbBlue
    
    lvwCallList.ColumnAdd 38, 0, 165, "�˓d��-(�K�X)"
    '��
    lvwCallList.SetColumnColor 38, vbBlue
    
    lvwCallList.ColumnAdd 39, 0, 165, "�˓d����-(�K�X)"
    '��
    lvwCallList.SetColumnColor 39, vbBlue
    
    lvwCallList.ColumnAdd 40, 0, 210, "����˓d����-(�K�X)"
    '��
    lvwCallList.SetColumnColor 40, vbBlue
    
    lvwCallList.ColumnAdd 41, 0, 165, "�s�����R-(�K�X)"
    '��
    lvwCallList.SetColumnColor 41, vbBlue
    
    lvwCallList.ColumnAdd 42, 0, 210, "�\������̓�-(�K�X)"
    '��
    lvwCallList.SetColumnColor 42, vbBlue
    
    
    lvwCallList.ColumnAdd 43, 0, 165, "�˓d����-(�ح����)"
    lvwCallList.ColumnAdd 44, 0, 165, "�˓d��-(�ح����)"
    lvwCallList.ColumnAdd 45, 0, 165, "�˓d����-(�ح����)"
    lvwCallList.ColumnAdd 46, 0, 230, "����˓d����-(�ح����)"
    lvwCallList.ColumnAdd 47, 0, 165, "�s�����R-(�ح����)"
    '��
    lvwCallList.SetColumnColor 43, RGB(91, 162, 97)
    '��
    lvwCallList.SetColumnColor 44, RGB(91, 162, 97)
    '��
    lvwCallList.SetColumnColor 45, RGB(91, 162, 97)
    '��
    lvwCallList.SetColumnColor 46, RGB(91, 162, 97)
    '��
    lvwCallList.SetColumnColor 47, RGB(91, 162, 97)
    
    
    
    lvwCallList.ColumnAdd 48, 0, 0, "ClientCode"
    lvwCallList.ColumnAdd 49, 0, 0, "CustomerCode"
    lvwCallList.ColumnAdd 50, 0, 0, "UserUniqueData"
    lvwCallList.ColumnAdd 51, 0, 0, "UserCategoryData"
    lvwCallList.ColumnAdd 52, 0, 0, "SystemPhoneNo1"
    lvwCallList.ColumnAdd 53, 0, 0, "SystemPhoneNo2"
    lvwCallList.ColumnAdd 54, 0, 0, "SystemPhoneNo3"
    lvwCallList.ColumnAdd 55, 0, 0, "CompleteFlg"

    'lvwCallList.SetFont "�l�r �o����", 11, 1, 0, 0

'�J������������
'Dim col As Long
'With lvwCallList
'   For col = 4 To .ColumnHeaders.Count - 1
''         SendMessage .hwnd, LVM_SETCOLUMNWIDTH, col, ByVal LVSCW_AUTOSIZE_USEHEADER
'   Next col
'End With
'�e�X�g�f�[�^
'Call mAddListItem
End Sub

'
'Private Sub mAddListItem()
'    Dim i As Integer
'    Dim ItemX As ListItem
'
' Do Until i >= 100
'        i = i + 1
'
'        Set ItemX = lvwCallList.ListItems.Add(i, , "xxxxxxxxxxx" & i)
'
'
'        ItemX.SubItems(1) = "" & "����" & i
'        ItemX.SubItems(2) = "" & "�g�E�L���E" & i
'        ItemX.SubItems(3) = "" & "�g�p�s��"
'        ItemX.SubItems(4) = "" & i & "�ĉ˓d"
'        ItemX.SubItems(5) = "" & i & "�_��"
'        ItemX.SubItems(6) = "" & i & "�_��X�V"
'        ItemX.SubItems(7) = "" & i & "�Z���ύX"
'        ItemX.SubItems(8) = "" & i & "���R�Ȃ�"
'        ItemX.SubItems(9) = "" & i & "����"
'
'        If (i - 1) Mod 3 = 0 Then
'            ItemX.SubItems(10) = "" & i & "�e"
'        Else
'        ItemX.SubItems(10) = "" & i & "�q"
'        End If
'
'        ItemX.SubItems(11) = "" & i & "DIS�g�E�L���E"
'        ItemX.SubItems(12) = "" & i & "���˗l"
'        ItemX.SubItems(13) = "" & i & "01000110101"
'        ItemX.SubItems(14) = "" & i & "6600110101"
'        ItemX.SubItems(15) = "" & i & "999-0010"
'        ItemX.SubItems(16) = "" & i
'        ItemX.SubItems(17) = "" & i
'        ItemX.SubItems(18) = "" & i
'        ItemX.SubItems(19) = "" & i
'        ItemX.SubItems(20) = "" & i
'        ItemX.SubItems(21) = "" & i
'        ItemX.SubItems(22) = "" & i
'        ItemX.SubItems(23) = "" & i
'        ItemX.SubItems(24) = "" & i
'        ItemX.SubItems(25) = "" & i
'        ItemX.SubItems(26) = "" & i
'        ItemX.SubItems(27) = "" & i
'        ItemX.SubItems(28) = "" & i
'        ItemX.SubItems(29) = "" & i
'        ItemX.SubItems(30) = "" & i
'        ItemX.SubItems(31) = "" & i
'        ItemX.SubItems(32) = "" & i
'        ItemX.SubItems(33) = "" & i
'        ItemX.SubItems(34) = "" & i
'        ItemX.SubItems(35) = "" & i
'        ItemX.SubItems(36) = "" & i
'        ItemX.SubItems(37) = "" & i
'        ItemX.SubItems(38) = "" & i
'        ItemX.SubItems(39) = "" & i
'        ItemX.SubItems(40) = "" & i
'        ItemX.SubItems(41) = "" & i
'        ItemX.SubItems(42) = "" & i
'        ItemX.SubItems(43) = "" & i
'        ItemX.SubItems(44) = "" & "-----"
'    Loop
'End Sub

'���ڑI�������ꍇ
Private Sub lvwCallList_OnKeyDown(VkCode As Long)
On Error GoTo Err_btnSelection_Click
    
    Dim rs As New ADODB.Recordset
    Dim SQLwhere As String
    Dim SQL As String
    Dim lCount As Long
    Dim strDialStartTimeFld As String
    Dim strDialEndTimeFld As String
    Dim i As Long
    Dim blnFlag As Boolean
    Dim ItemX As ListItem
    Dim subWhere As String
    Dim str As String
    Dim vntData     As Variant
    Dim iItem As Long
        
    '�I����Ԃ�����
    iItem = lvwCallList.ItemSelect
    If iItem = -1 Then
        Exit Sub
    End If
    
    Select Case VkCode
    Case 1      '���N���b�N
        Exit Sub
    Case 2      '�E�N���b�N
        Exit Sub
    Case 4      '���v�N���b�N
        '��������
        'Exit Sub
    Case 13     'VK_RETURN
        '��������
        'Exit Sub
    Case Else   '���̑��L�[(���z�L�[�R�[�h)
        Exit Sub
    End Select
    
    '�I���s�擾
    str = lvwCallList.ItemGet(iItem)
    '�^�u�ŕ���
    vntData = Split(str, "" & vbTab)
        
    If UBound(vntData) = 0 Then Exit Sub
    
    gQuestionAnswer = False
    
    '�������� �J�X�^�}�R�[�h
'    SQLwhere = "WHERE (((TB_DIALCTL.ClientCode)=" & gLogonInfo.ClientCode & ") " & _
'                "AND ((TB_DIALCTL.CustomerCode)='" & vntData(48).Text & "')) "
                
    SQLwhere = "WHERE (((TB_DIALCTL.ClientCode)=" & gLogonInfo.ClientCode & ") " & _
                "AND ((TB_DIALCTL.CustomerCode)='" & vntData(49) & "')) "
                
                '"AND ((TB_DIALCTL.CompleteFlag)='0') " & _
                '"AND ((TB_CLIENTDATA.StopFlag)='0'))" & SQLwhere

    '�˓d���X�g���擾
    SQL = "SELECT Count(TB_DIALCTL.ProfileID) AS DataCount " & _
            "FROM (TB_DIALCTL LEFT JOIN TB_CLIENTDATA ON (TB_DIALCTL.CustomerCode = TB_CLIENTDATA.CustomerCode) AND (TB_DIALCTL.ClientCode = TB_CLIENTDATA.ClientCode)) LEFT JOIN TB_CLIENTSUBDATA ON (TB_DIALCTL.CustomerCode = TB_CLIENTSUBDATA.CustomerCode) AND (TB_DIALCTL.ClientCode = TB_CLIENTSUBDATA.ClientCode) " & _
            "" & SQLwhere
    
    ConnectionOpen gDBString
    
    rs.Open SQL, gCnn
    lCount = rs!DataCount
    rs.Close
    
    '���X�g���Ȃ��ꍇ�A�����́A�q�̏ꍇ�AINBound�Ƃ��ď�������
    If lCount = 0 Or vntData(12) = "�q" Then
        '�@�l��
        gKeyInfo.CustomerName = vntData(13) & ""
        '���q�l�R�[�h
'        gKeyInfo.CustomerCode = vntData(48)
        gKeyInfo.CustomerCode = vntData(49) & ""
        gKeyInfo.ReloadFlag = True      '
        '���́u���C���v���d�b�����ƌ��т��Ă���Ƃ��̂ݘA������
        If frmMain.ucPhoneApi1.MainIndex = CStr(frmMain.hwnd) Then
'            frmMain.ucPhoneApi1.PhoneNo = vntData(51)    '�d�b�ԍ���d�b�@�ɒm�点��
            frmMain.ucPhoneApi1.PhoneNo = vntData(52) & ""    '�d�b�ԍ���d�b�@�ɒm�点��
        End If
        'IBound�\��
        ISIBFLG = True
        gQuestionAnswer = True
        Me.Visible = False
    Else
        'OutBound�̏ꍇ
        SQL = "SELECT TB_DIALCTL.ProfileID,TB_DIALCTL.CustomerCode,TB_DIALCTL.dialcount,TB_DIALCTL.CompleteFlag " & _
        "FROM (TB_DIALCTL LEFT JOIN TB_CLIENTDATA ON (TB_DIALCTL.CustomerCode = TB_CLIENTDATA.CustomerCode) AND (TB_DIALCTL.ClientCode = TB_CLIENTDATA.ClientCode)) LEFT JOIN TB_CLIENTSUBDATA ON (TB_DIALCTL.CustomerCode = TB_CLIENTSUBDATA.CustomerCode) AND (TB_DIALCTL.ClientCode = TB_CLIENTSUBDATA.ClientCode) " & _
        "" & SQLwhere & " order by TB_DIALCTL.dialcount, TB_DIALCTL.LastDialDate, TB_DIALCTL.TalkCount, TB_DIALCTL.LastTalkDate, TB_DIALCTL.ProfileID"


        rs.Open SQL, gCnn
        '���M���ر
         Erase gListCallData
         subWhere = ""
         Do Until rs.EOF
            i = i + 1
            '���M���\��
            ReDim Preserve gListCallData(i) As uListCallData
            gListCallData(i).ProfileID = rs!ProfileID & ""
            gListCallData(i).CustomerCode = rs!CustomerCode & ""
            gListCallData(i).DialCount = rs!DialCount & ""
            
            'ProfileID�����ŏ��擾
            subWhere = subWhere & "'" & rs!ProfileID & "',"
            rs.MoveNext
        Loop
        rs.Close

        '�˓d���X�g���擾
         If Right(subWhere, 1) = "," Then
            '�Ō�̃J���}�Ǝ�菜��
            subWhere = Left(subWhere, Len(subWhere) - 1)
         End If
         If subWhere <> "" Then
            'TB_DIALCTL.ProfileID �ŏ��擾
            subWhere = " WHERE " & "TB_DIALCTL.ProfileID IN (" & subWhere & ")"
            frmCallListSelect.getCallListInfo (subWhere)
         End If
         
         '�@�l��
        gKeyInfo.CustomerName = vntData(13) & ""
        '���q�l�R�[�h
'        gKeyInfo.CustomerCode = vntData(48) & ""
        gKeyInfo.CustomerCode = vntData(49) & ""
        
         frmCallListSelect.Show (1)
         '���X�g�I���ς݃��C����ʔ��
         If selctStateYesNo = True Then
            '��\��
            Me.Visible = False
            gQuestionAnswer = True
         End If
    End If
Exit_btnSelection_Click:
'    prgBar.Value = 0
    
'    Me.Visible = False 'Unload Me�̌�ɃE�C���h�E���̃R���g���[����M��ƃ������Ɏc���Ă��܂��̂�Unload�Ȍ�ɏ���������Ȃ�
    Exit Sub

Err_btnSelection_Click:
    MsgBox Err.Description
    Resume Exit_btnSelection_Click

End Sub


'------------------------------------------------------------------
'�v���V�[�W����:  :txtUserPass_KeyPress
'����:            :�˓d�󋵃L�[�{�[�h����
'����:            :KeyAscii           O   xxxx
'------------------------------------------------------------------
Private Sub txtElectric1State_KeyPress(KeyAscii As Integer)
    Select Case KeyAscii
    Case vbKeyReturn
        If chkElectric1State.Value <> 0 Then
            txtElectric1State.SetFocus
        End If
    Case Else
         KeyAscii = vbNull
        
    End Select
End Sub

'------------------------------------------------------------------
'�v���V�[�W����:  :txtElectric2State_KeyPress
'����:            :�˓d�󋵃L�[�{�[�h����
'����:            :KeyAscii           O   xxxx
'------------------------------------------------------------------
Private Sub txtElectric2State_KeyPress(KeyAscii As Integer)
    Select Case KeyAscii
    Case vbKeyReturn
        If chkElectric2State.Value <> 0 Then
            txtElectric2State.SetFocus
        End If
    Case Else
         KeyAscii = vbNull
        
    End Select
End Sub


'------------------------------------------------------------------
'�v���V�[�W����:  :txtGasState_KeyPress
'����:            :�˓d�󋵃L�[�{�[�h����
'����:            :KeyAscii           O   xxxx
'------------------------------------------------------------------
Private Sub txtGasState_KeyPress(KeyAscii As Integer)
    Select Case KeyAscii
    Case vbKeyReturn
        If chkGasState.Value <> 0 Then
            txtGasState.SetFocus
        End If
    Case Else
         KeyAscii = vbNull
        
    End Select
End Sub


'------------------------------------------------------------------
'�v���V�[�W����:  :txtSolutionState_KeyPress
'����:            :�˓d�󋵃L�[�{�[�h����
'����:            :KeyAscii           O   xxxx
'------------------------------------------------------------------
Private Sub txtSolutionState_KeyPress(KeyAscii As Integer)
    Select Case KeyAscii
    Case vbKeyReturn
        If chkSolutionState.Value <> 0 Then
            txtSolutionState.SetFocus
        End If
    Case Else
'        KeyAscii = f_InputCheck_Textbox_X(Me.txtUserPass, KeyAscii, 15)
            KeyAscii = vbNull
        
    End Select
End Sub


'------------------------------------------------------------------
'�v���V�[�W����:  :txtStaffLocation_KeyPress
'����:            :�˓d�󋵃L�[�{�[�h����
'����:            :KeyAscii           O   xxxx
'------------------------------------------------------------------
Private Sub txtStaffLocation_KeyPress(KeyAscii As Integer)
    Select Case KeyAscii
    Case vbKeyReturn
        If chkStaffLocation.Value <> 0 Then
            txtStaffLocation.SetFocus
        End If
    Case Else
'        KeyAscii = f_InputCheck_Textbox_X(Me.txtUserPass, KeyAscii, 15)
            KeyAscii = vbNull
        
    End Select
End Sub




'******************************************************************
'�v���V�[�W����:  :getDialin
'����:            :���M�ԍ��擾
'����:            :ListCode           I   xxxx
'                 :JobCode            I   xxxx
'�߂�l           :String�^
'******************************************************************
Function getDialin(ByVal ListCode As String, ByVal JobCode As String) As String

    Dim cnn As New ADODB.Connection
    Dim rs  As New ADODB.Recordset
    Dim SQL As String
    Dim BusyResultName As String
    Dim dailKey As Variant
    Dim listDialin As String: listDialin = ""
    Dim freeNo As String: freeNo = ""
    Dim popUpName As String: popUpName = ""
    Dim j As Integer

On Error GoTo Err_getDialin:

    cnn.Open gDBString
    
'    SQL = "select DialinNo,FreeDial from tb_dialin where ClientCode=" & gLogonInfo.ClientCode & " order by DialinNo"
    SQL = "select dialinNo,ListName,aliasname,jobcode from tb_list where listCode='" & ListCode & "' And JobCode='" & JobCode & "'"
    
    rs.Open SQL, cnn
    Do Until rs.EOF
         listDialin = rs.Fields("DialinNo").Value & ""
        rs.MoveNext
    Loop
    rs.Close
  
    
    If listDialin <> "" Then
        SQL = "select * from tb_dialin where dialinNo = '" & listDialin & "'"
        rs.Open SQL, cnn
         Do Until rs.EOF
            freeNo = rs.Fields("FreeDial").Value & ""
            popUpName = rs.Fields("PopupName").Value & ""
            rs.MoveNext
         Loop
    End If
    
    For j = 0 To frmMain.cobDialNumber.ListCount
       If frmMain.cobDialNumber.List(j) = popUpName Then
            frmMain.cobDialNumber.ListIndex = j
            frmMain.ucPhoneApi1.CallingNumber = getFreeDialByPopName(popUpName)
        Exit For
       End If
    Next j
    
    getDialin = freeNo
    
    
     cnn.Close
    
'    For Each dailKey In dc.Keys
'        cobDialNumber.AddItem dcPopName.item(dailKey)           'cobDialNumber.ItemData(cobDialNumber.NewIndex) = "" & rs!FreeDial       '�o�^"
'    Next
'
'    '���ڑ��݃`�F�b�N
'    If cobDialNumber.ListCount = 0 Then
'        Exit Function
'    End If
'
    '�R���{�{�b�N�X���ڂ͉����I�����Ă��Ȃ����
'    cobDialNumber.ListIndex = -1
'
'
'    Dim i As Integer
'    '�����I������
'    If dcPopName.Exists(listDialin) Then
'        For i = 0 To cobDialNumber.ListCount
'            If cobDialNumber.List(i) = dcPopName.item(listDialin) Then
'                cobDialNumber.ListIndex = i
'                Exit For
'            End If
'        Next i
'    End If

    Exit Function
Err_getDialin:
    MsgBox (Err.Description)
    Exit Function
End Function



'------------------------------------------------------------------
'�v���V�[�W����:  :lvwCallList_OnGetData
'����:            :�������ʐݒ�
'����:            :iItem              O   xxxx
'                 :iSubItem           O   xxxx
'                 :rs                 O   xxxx
'                 :strOut             O   xxxx
'------------------------------------------------------------------
Private Sub lvwCallList_OnGetData(iItem As Long, iSubItem As Long, rs As ADODB.Recordset, strOut As String)
'    If Not OnceFlg Then
'        '�f�[�^����
'        lblDataCount.Caption = Replace(lblDataCount.Caption, "0", CommaSplitFmt(rs.Fields("RowCount")))
'        'lblErrMsg �����`�F�b�N
'        If CLng(rs.Fields("RowCount")) >= MAX_ROW_NUM Then
'            lblErrMsg.Visible = True
'        End If
'        OnceFlg = True
'    End If
    
    Select Case iSubItem
      Case 0
        strOut = ""
      Case 1
        '���q���ܔԍ�
        strOut = "" & rs.Fields("Data001").Value
      Case 2
        '�n�_��
        strOut = "" & rs.Fields("Data041").Value
      Case 3
        '�n�_���J�i
        strOut = "" & rs.Fields("Data042").Value
      Case 4
       '�g�p���
        'ItemX.SubItems(4) = "" & rs.Fields("Data010")
        strOut = convertMethod.mConvertValue2Name("" & rs.Fields("Data010").Value, "compcrm_shiyo_jotai")
      Case 5
        '�����Ώے������R
        strOut = "" & rs.Fields("Data044").Value
      Case 6
        '��ĒS���ӏ� �R�[�h����l�ύX
        'ItemX.SubItems(6) = "" & rs.Fields("Data009")
        'ItemX.SubItems(6) = convertMethod.mConvertValue2Code("" & rs.Fields("Data009"))
        strOut = convertMethod.mConvertValue2Name("" & rs.Fields("Data009").Value, "compcrm_teian_tanto_kasho")
      Case 7
       '�˓d�ړI�d�C�P
        strOut = "" & rs.Fields("Data049").Value
      Case 8
        '�˓d�ړI�d�C2
        strOut = "" & rs.Fields("Data050").Value
      Case 9
        '�˓d�ړI�K�X
        strOut = "" & rs.Fields("Data051").Value
      Case 10
        '�˓d�ړI�\�����[�V����
        strOut = "" & rs.Fields("Data052").Value
      Case 11
       '�Ǝ�
        strOut = "" & rs.Fields("Data045").Value
      Case 12
        '�e�q
        'ItemX.SubItems(12) = "" & rs.Fields("Data046")
        'strOut = "" & rs.Fields("ParentChildJP").Value
        'START 20210924 �e�q���肪""���e�Ƃ��ď�������B
        If Trim("" & rs.Fields("Data046").Value) = "" Then
            strOut = ""
        Else
            strOut = "" & rs.Fields("ParentChildJP").Value
        End If
        'END 20210924 �e�q���肪""���e�Ƃ��ď�������B
      Case 13
        '�@�l��
        strOut = "" & rs.Fields("Data043").Value
      Case 14
        '�˓d�S����
        strOut = "" & rs.Fields("Data098").Value
      Case 15
        '��\TEL
        strOut = "" & rs.Fields("Data002").Value
      Case 16
        '�x����TEL
        strOut = "" & rs.Fields("Data004").Value
      Case 17
        '���v�ꏊTEL
        strOut = "" & rs.Fields("Data003").Value
      Case 18
        '���v�ꏊ�Z��
        strOut = "" & rs.Fields("Data096").Value
      Case 19
        '�x����TEL
        strOut = "" & rs.Fields("Data005").Value
      Case 20
        '�x����Z��
        strOut = "" & rs.Fields("Data097").Value
      Case 21
        '�x���Җ�
        strOut = "" & rs.Fields("Data082").Value
      Case 22
        '�ĉ˓d��
        strOut = "" & rs.Fields("ReCallDate").Value
      Case 23
        '�˓d���ԑ�(�n)
        strOut = "" & rs.Fields("CallStartTime").Value
      Case 24
        '�˓d���ԑ�(�I)
        strOut = "" & rs.Fields("CallEndTime").Value
      Case 25
        '�˓d�����i�d�C1�j
        strOut = "" & rs.Fields("Den1CallDate").Value
      Case 26
        '�˓d�󋵁@�d�C�P
        strOut = "" & ConvertCode2Value("" & rs.Fields("Den1CallState").Value)
      Case 27
        '�˓d���ʁ@�d�C�P
        strOut = "" & ConvertCode2Value("" & rs.Fields("Den1CallResult").Value)
      Case 28
        '����˓d�����d�C�P
        strOut = "" & rs.Fields("Den1NextCallDate").Value
      Case 29
        '�s�����R
        strOut = "" & ConvertCode2Value("" & rs.Fields("Den1FailReason").Value)
      Case 30
        'E�\������̓�(�d�P)
        strOut = "" & rs.Fields("Data027").Value
      Case 31
        ''�˓d�����i�d�C2�j
        strOut = "" & rs.Fields("Den2CallDate").Value
      Case 32
        '�˓d�� �d�C2
        strOut = "" & ConvertCode2Value("" & rs.Fields("Den2CallState").Value)
      Case 33
        '�˓d���ʓd�C2
        strOut = "" & ConvertCode2Value("" & rs.Fields("Den2CallResult").Value)
      Case 34
        '����˓d�����d�C2
        strOut = "" & rs.Fields("Den2NextCallDate").Value
      Case 35
        '�s�����R
        strOut = "" & ConvertCode2Value("" & rs.Fields("Den2FailReason").Value)
      Case 36
        'E2�\������̓�
        'ItemX.SubItems(36) = "" & rs.Fields("Den2ApplyDate")
        strOut = "" & rs.Fields("Data039").Value
      Case 37
        '�˓d����
        strOut = "" & rs.Fields("GasCallDate").Value
      Case 38
        '�˓d�� �K�X
        strOut = "" & ConvertCode2Value("" & rs.Fields("GasCallState").Value)
      Case 39
        '�˓d���ʃK�X
        strOut = "" & ConvertCode2Value("" & rs.Fields("GasCallResult").Value)
      Case 40
        '����˓d����
        strOut = "" & rs.Fields("GasNextCallDate").Value
      Case 41
        '�s�����R
        strOut = "" & ConvertCode2Value("" & rs.Fields("GasFailReason").Value)
      Case 42
        'G�\������̓�
        strOut = "" & rs.Fields("Data038").Value
      Case 43
        '�˓d�����\�����[�V����
        strOut = "" & rs.Fields("SoluCallDate").Value
      Case 44
        '�˓d�� �ح����
        strOut = "" & ConvertCode2Value("" & rs.Fields("SoluCallState").Value)
      Case 45
        strOut = "" & ConvertCode2Value("" & rs.Fields("SoluCallResult").Value)
      Case 46
        strOut = "" & rs.Fields("SoluNextCallDate").Value
      Case 47
        '�s�����R
        strOut = "" & ConvertCode2Value("" & rs.Fields("SoluFailReason").Value)
      Case 48
        '�N���C�A���g�R�[�h
        strOut = "" & rs.Fields("ClientCode").Value
      Case 49
        '�J�X�^�}�R�[�h
        strOut = "" & rs.Fields("CustomerCode").Value
      Case 50
        strOut = "" & rs.Fields("UserUniqueData").Value
      Case 51
        strOut = "" & rs.Fields("UserCategoryData").Value
      Case 52
        '�d�b�ԍ�
        strOut = "" & rs.Fields("SystemPhoneNo1").Value
      Case 53
        '�d�b�ԍ�2
        strOut = "" & rs.Fields("SystemPhoneNo2").Value
      Case 54
        '�d�b�ԍ�3
        strOut = "" & rs.Fields("SystemPhoneNo3").Value
      Case 55
        '����FLG
        strOut = "" & rs.Fields("completeFlg").Value
        '�����̏ꍇ��ڰ�ŕ\��
        If "" & rs.Fields("completeFlg").Value = "1" Then
            '��ڰ�\��
            lvwCallList.SetItemBackColor iItem, COLOR_GRAY
        Else
            '�e�̏ꍇ�I�����W�ŕ\��
            If "" & rs.Fields("Data046").Value = "0" Or Trim("" & rs.Fields("Data046").Value) = "" Then
                '�d�C1�˓d�ړI����̏ꍇ�A�I�����W
                If "" & rs.Fields("Data049").Value <> "" Then
                    lvwCallList.SetItemBackColor iItem, COLOR_ORANGE
                End If
                
                '�d�C2�˓d�ړI����̏ꍇ�A�I�����W
                If "" & rs.Fields("Data050").Value <> "" Then
                    lvwCallList.SetItemBackColor iItem, COLOR_ORANGE
                End If
                
                '�K�X�˓d�ړI����̏ꍇ�A�I�����W
                If "" & rs.Fields("Data051").Value <> "" Then
                    lvwCallList.SetItemBackColor iItem, COLOR_ORANGE
                End If
                
                '�\�����[�V�����˓d�ړI����̏ꍇ�A�I�����W
                If "" & rs.Fields("Data052").Value <> "" Then
                    lvwCallList.SetItemBackColor iItem, COLOR_ORANGE
                End If
            'START 20220601 �ǉ��v�]A �w�i�F�ύX
            Else
                              '�d�C1�˓d�ړI����̏ꍇ�A�I�����W
                If "" & rs.Fields("Data049").Value <> "" Then
                    lvwCallList.SetItemBackColor iItem, COLOR_THIN_ORANGE
                End If
                
                '�d�C2�˓d�ړI����̏ꍇ�A�I�����W
                If "" & rs.Fields("Data050").Value <> "" Then
                    lvwCallList.SetItemBackColor iItem, COLOR_THIN_ORANGE
                End If
                
                '�K�X�˓d�ړI����̏ꍇ�A�I�����W
                If "" & rs.Fields("Data051").Value <> "" Then
                    lvwCallList.SetItemBackColor iItem, COLOR_THIN_ORANGE
                End If
                
                '�\�����[�V�����˓d�ړI����̏ꍇ�A�I�����W
                If "" & rs.Fields("Data052").Value <> "" Then
                    lvwCallList.SetItemBackColor iItem, COLOR_THIN_ORANGE
                End If
            'END 20220601 �ǉ��v�]A �w�i�F�ύX
            End If
        End If
      Case Else
          strOut = ""
      End Select
End Sub

