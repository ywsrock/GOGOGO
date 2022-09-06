VERSION 5.00
Object = "{831FDD16-0C5C-11D2-A9FC-0000F8754DA1}#2.2#0"; "MSCOMCTL.OCX"
Object = "{F9043C88-F6F2-101A-A3C9-08002B2F49FB}#1.2#0"; "COMDLG32.OCX"
Object = "{42FFC9A8-7FC5-4D28-95E0-9273BAC25E7D}#3.0#0"; "lvwODDB2.ocx"
Begin VB.Form frmNewListCall 
   BackColor       =   &H00008000&
   BorderStyle     =   1  '固定(実線)
   Caption         =   "架電リスト一覧"
   ClientHeight    =   12030
   ClientLeft      =   45
   ClientTop       =   1440
   ClientWidth     =   17700
   ControlBox      =   0   'False
   BeginProperty Font 
      Name            =   "ＭＳ Ｐ明朝"
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
         Caption         =   "実行モード"
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
            Caption         =   "検索"
            BeginProperty Font 
               Name            =   "ＭＳ Ｐ明朝"
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
            Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
            TabIndex        =   56
            Top             =   960
            Width           =   1035
         End
         Begin VB.CommandButton btnExport 
            Caption         =   "出力"
            BeginProperty Font 
               Name            =   "ＭＳ Ｐ明朝"
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
            Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
            TabIndex        =   57
            Top             =   960
            Width           =   1035
         End
      End
      Begin VB.Frame frmReadCheck 
         BackColor       =   &H00C0E0FF&
         Caption         =   "■検索条件"
         Height          =   3735
         Left            =   240
         TabIndex        =   59
         Top             =   120
         Width           =   14355
         Begin VB.CommandButton btnStaffLocation 
            Caption         =   "★"
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
               Caption         =   "提案担当箇所"
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
            Caption         =   "★"
            Enabled         =   0   'False
            Height          =   345
            Left            =   13815
            TabIndex        =   33
            Top             =   1080
            Width           =   345
         End
         Begin VB.CommandButton btnNextCallTimeStart 
            Caption         =   "★"
            Enabled         =   0   'False
            Height          =   345
            Left            =   4560
            TabIndex        =   24
            Top             =   3240
            Width           =   345
         End
         Begin VB.CommandButton btnNextCallTimeEnd 
            Caption         =   "★"
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
            IMEMode         =   3  'ｵﾌ固定
            Left            =   5265
            TabIndex        =   25
            TabStop         =   0   'False
            Top             =   3240
            Width           =   1635
         End
         Begin VB.TextBox Text13 
            Alignment       =   2  '中央揃え
            BackColor       =   &H8000000F&
            BeginProperty Font 
               Name            =   "ＭＳ Ｐ明朝"
               Size            =   12
               Charset         =   128
               Weight          =   700
               Underline       =   0   'False
               Italic          =   0   'False
               Strikethrough   =   0   'False
            EndProperty
            Height          =   360
            IMEMode         =   3  'ｵﾌ固定
            Left            =   4800
            TabIndex        =   88
            TabStop         =   0   'False
            Text            =   "～"
            Top             =   3240
            Width           =   555
         End
         Begin VB.TextBox txtNextCallTimeStart 
            BackColor       =   &H8000000F&
            Height          =   345
            IMEMode         =   3  'ｵﾌ固定
            Left            =   3060
            TabIndex        =   23
            TabStop         =   0   'False
            Top             =   3240
            Width           =   1635
         End
         Begin VB.CommandButton btnReCallDateStart 
            Caption         =   "★"
            Enabled         =   0   'False
            Height          =   345
            Left            =   4560
            TabIndex        =   19
            Top             =   2880
            Width           =   345
         End
         Begin VB.CommandButton btnReCallDateEnd 
            Caption         =   "★"
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
            IMEMode         =   3  'ｵﾌ固定
            Left            =   5265
            TabIndex        =   20
            TabStop         =   0   'False
            Top             =   2880
            Width           =   1635
         End
         Begin VB.TextBox Text10 
            Alignment       =   2  '中央揃え
            BackColor       =   &H8000000F&
            BeginProperty Font 
               Name            =   "ＭＳ Ｐ明朝"
               Size            =   12
               Charset         =   128
               Weight          =   700
               Underline       =   0   'False
               Italic          =   0   'False
               Strikethrough   =   0   'False
            EndProperty
            Height          =   345
            IMEMode         =   3  'ｵﾌ固定
            Left            =   4800
            TabIndex        =   87
            TabStop         =   0   'False
            Text            =   "～"
            Top             =   2880
            Width           =   555
         End
         Begin VB.CommandButton btnSolutionState 
            Caption         =   "★"
            Enabled         =   0   'False
            Height          =   345
            Left            =   13815
            TabIndex        =   42
            Top             =   2160
            Width           =   345
         End
         Begin VB.CommandButton btnGasState 
            Caption         =   "★"
            Enabled         =   0   'False
            Height          =   345
            Left            =   13815
            TabIndex        =   39
            Top             =   1800
            Width           =   345
         End
         Begin VB.CommandButton btnElectric2State 
            Caption         =   "★"
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
            IMEMode         =   3  'ｵﾌ固定
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
            IMEMode         =   3  'ｵﾌ固定
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
            IMEMode         =   3  'ｵﾌ固定
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
            IMEMode         =   3  'ｵﾌ固定
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
               Caption         =   "架電担当者"
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
               Caption         =   "架電目的（電気1）"
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
               Caption         =   "架電目的（電気2）"
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
               Caption         =   "架電目的（ガス）"
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
               Caption         =   "架電目的（ｿﾘｭｰｼｮﾝ）"
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
               Caption         =   "電話番号"
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
               Caption         =   "再架電予定日"
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
               Caption         =   "次回架電日時"
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
               Caption         =   "企業名"
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
               Caption         =   "法人名"
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
               Caption         =   "架電状況（電気1）"
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
               Caption         =   "架電状況（電気2）"
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
               Caption         =   "架電状況（ガス）"
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
               Caption         =   "架電状況（ｿﾘｭｰｼｮﾝ）"
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
               Caption         =   "申込書（電気1）"
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
               Caption         =   "申込書（電気2）"
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
               Caption         =   "受領待"
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
               Caption         =   "受領済"
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
               Caption         =   "受領不要"
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
               Caption         =   "受領待"
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
               Caption         =   "受領済"
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
               Caption         =   "受領不要"
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
            Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
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
            Text            =   "顧客名"
            Object.Width           =   3351
         EndProperty
         BeginProperty ColumnHeader(3) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            SubItemIndex    =   2
            Text            =   "電話番号"
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
            Text            =   "優先結果"
            Object.Width           =   3351
         EndProperty
         BeginProperty ColumnHeader(7) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            Alignment       =   2
            SubItemIndex    =   6
            Text            =   "予定日(平日)"
            Object.Width           =   2822
         EndProperty
         BeginProperty ColumnHeader(8) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            Alignment       =   2
            SubItemIndex    =   7
            Text            =   "開始(平日)"
            Object.Width           =   2469
         EndProperty
         BeginProperty ColumnHeader(9) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            Alignment       =   2
            SubItemIndex    =   8
            Text            =   "終了(平日)"
            Object.Width           =   2469
         EndProperty
         BeginProperty ColumnHeader(10) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            Alignment       =   2
            SubItemIndex    =   9
            Text            =   "予定日(休日)"
            Object.Width           =   0
         EndProperty
         BeginProperty ColumnHeader(11) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            Alignment       =   2
            SubItemIndex    =   10
            Text            =   "開始(休日)"
            Object.Width           =   0
         EndProperty
         BeginProperty ColumnHeader(12) {BDD1F052-858B-11D1-B16A-00C0F0283628} 
            Alignment       =   2
            SubItemIndex    =   11
            Text            =   "終了(休日)"
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
         Alignment       =   1  '右揃え
         BackStyle       =   0  '透明
         Caption         =   "エラー：検索結果件数が多すぎて結果が表示できません、検索条件を絞ってください。"
         ForeColor       =   &H000000C0&
         Height          =   255
         Left            =   -4560
         TabIndex        =   86
         Top             =   4080
         Visible         =   0   'False
         Width           =   13275
      End
      Begin VB.Label lblDataCount 
         Alignment       =   1  '右揃え
         BackStyle       =   0  '透明
         Caption         =   "0件のデータが該当しました"
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
