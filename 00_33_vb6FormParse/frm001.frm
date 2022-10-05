VERSION 5.00
Begin VB.Form frmMain 
   BackColor       =   &H00008000&
   BorderStyle     =   3  'å≈íË¿ﬁ≤±€∏ﬁ
   ClientHeight    =   13455
   ClientLeft      =   45
   ClientTop       =   45
   ClientWidth     =   1470
   ControlBox      =   0   'False
   BeginProperty Font 
      Name            =   "ÇlÇr Çoñæí©"
      Size            =   11.25
      Charset         =   128
      Weight          =   700
      Underline       =   0   'False
      Italic          =   0   'False
      Strikethrough   =   0   'False
   EndProperty
   LinkTopic       =   "Form1"
   LockControls    =   -1  'True
   MaxButton       =   0   'False
   MinButton       =   0   'False
   ScaleHeight     =   13455
   ScaleWidth      =   1470
   ShowInTaskbar   =   0   'False
   Begin PhoneAPI.ucPhoneApi ucPhoneApi1 
      Height          =   1215
      Left            =   1080
      TabIndex        =   14
      Top             =   660
      Visible         =   0   'False
      Width           =   1275
      _ExtentX        =   2249
      _ExtentY        =   2143
   End
   Begin afValue.Value afValue 
      Left            =   1080
      Top             =   0
      _ExtentX        =   635
      _ExtentY        =   344
   End
   Begin VB.PictureBox Picture1 
      BackColor       =   &H00C0E0FF&
      Height          =   13335
      Left            =   60
      ScaleHeight     =   13275
      ScaleWidth      =   1275
      TabIndex        =   0
      Top             =   60
      Width           =   1335
      Begin VB.CommandButton btnMenu 
         Caption         =   "ÉVÉXÉeÉÄ"
         BeginProperty Font 
            Name            =   "ÇlÇr Çoñæí©"
            Size            =   9.75
            Charset         =   128
            Weight          =   700
            Underline       =   0   'False
            Italic          =   0   'False
            Strikethrough   =   0   'False
         EndProperty
         Height          =   900
         Index           =   12
         Left            =   120
         Picture         =   "frmMain.frx":0000
         Style           =   1  '∏ﬁ◊Ã®Ø∏Ω
         TabIndex        =   13
         Top             =   12300
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "É}ÉXÉ^"
         BeginProperty Font 
            Name            =   "ÇlÇr Çoñæí©"
            Size            =   9.75
            Charset         =   128
            Weight          =   700
            Underline       =   0   'False
            Italic          =   0   'False
            Strikethrough   =   0   'False
         EndProperty
         Height          =   900
         Index           =   11
         Left            =   120
         Picture         =   "frmMain.frx":030A
         Style           =   1  '∏ﬁ◊Ã®Ø∏Ω
         TabIndex        =   12
         Top             =   11285
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "ÉÜÅ[ÉU"
         BeginProperty Font 
            Name            =   "ÇlÇr Çoñæí©"
            Size            =   9.75
            Charset         =   128
            Weight          =   700
            Underline       =   0   'False
            Italic          =   0   'False
            Strikethrough   =   0   'False
         EndProperty
         Height          =   900
         Index           =   10
         Left            =   120
         Picture         =   "frmMain.frx":0614
         Style           =   1  '∏ﬁ◊Ã®Ø∏Ω
         TabIndex        =   11
         Top             =   10270
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "êÍópã∆ñ±"
         BeginProperty Font 
            Name            =   "ÇlÇr Çoñæí©"
            Size            =   9.75
            Charset         =   128
            Weight          =   700
            Underline       =   0   'False
            Italic          =   0   'False
            Strikethrough   =   0   'False
         EndProperty
         Height          =   900
         Index           =   9
         Left            =   120
         Picture         =   "frmMain.frx":091E
         Style           =   1  '∏ﬁ◊Ã®Ø∏Ω
         TabIndex        =   10
         Top             =   9255
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "CTI"
         BeginProperty Font 
            Name            =   "ÇlÇr Çoñæí©"
            Size            =   9.75
            Charset         =   128
            Weight          =   700
            Underline       =   0   'False
            Italic          =   0   'False
            Strikethrough   =   0   'False
         EndProperty
         Height          =   900
         Index           =   8
         Left            =   120
         Picture         =   "frmMain.frx":0C28
         Style           =   1  '∏ﬁ◊Ã®Ø∏Ω
         TabIndex        =   9
         Top             =   8240
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "ä«óù"
         BeginProperty Font 
            Name            =   "ÇlÇr Çoñæí©"
            Size            =   9.75
            Charset         =   128
            Weight          =   700
            Underline       =   0   'False
            Italic          =   0   'False
            Strikethrough   =   0   'False
         EndProperty
         Height          =   900
         Index           =   7
         Left            =   120
         Picture         =   "frmMain.frx":0F32
         Style           =   1  '∏ﬁ◊Ã®Ø∏Ω
         TabIndex        =   8
         Top             =   7225
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "ÉAÉCÉeÉÄ"
         BeginProperty Font 
            Name            =   "ÇlÇr Çoñæí©"
            Size            =   9.75
            Charset         =   128
            Weight          =   700
            Underline       =   0   'False
            Italic          =   0   'False
            Strikethrough   =   0   'False
         EndProperty
         Height          =   900
         Index           =   6
         Left            =   120
         Picture         =   "frmMain.frx":123C
         Style           =   1  '∏ﬁ◊Ã®Ø∏Ω
         TabIndex        =   7
         Top             =   6210
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "ÉfÅ[É^"
         BeginProperty Font 
            Name            =   "ÇlÇr Çoñæí©"
            Size            =   9.75
            Charset         =   128
            Weight          =   700
            Underline       =   0   'False
            Italic          =   0   'False
            Strikethrough   =   0   'False
         EndProperty
         Height          =   900
         Index           =   5
         Left            =   120
         Picture         =   "frmMain.frx":1546
         Style           =   1  '∏ﬁ◊Ã®Ø∏Ω
         TabIndex        =   6
         Top             =   5195
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "í†ï["
         BeginProperty Font 
            Name            =   "ÇlÇr Çoñæí©"
            Size            =   9.75
            Charset         =   128
            Weight          =   700
            Underline       =   0   'False
            Italic          =   0   'False
            Strikethrough   =   0   'False
         EndProperty
         Height          =   900
         Index           =   4
         Left            =   120
         Picture         =   "frmMain.frx":1850
         Style           =   1  '∏ﬁ◊Ã®Ø∏Ω
         TabIndex        =   5
         Top             =   4180
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "ê\ëó"
         BeginProperty Font 
            Name            =   "ÇlÇr Çoñæí©"
            Size            =   9.75
            Charset         =   128
            Weight          =   700
            Underline       =   0   'False
            Italic          =   0   'False
            Strikethrough   =   0   'False
         EndProperty
         Height          =   900
         Index           =   3
         Left            =   120
         Picture         =   "frmMain.frx":1B5A
         Style           =   1  '∏ﬁ◊Ã®Ø∏Ω
         TabIndex        =   4
         Top             =   3165
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "àÀóä"
         BeginProperty Font 
            Name            =   "ÇlÇr Çoñæí©"
            Size            =   9.75
            Charset         =   128
            Weight          =   700
            Underline       =   0   'False
            Italic          =   0   'False
            Strikethrough   =   0   'False
         EndProperty
         Height          =   900
         Index           =   2
         Left            =   120
         Picture         =   "frmMain.frx":1E64
         Style           =   1  '∏ﬁ◊Ã®Ø∏Ω
         TabIndex        =   3
         Top             =   2145
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   ""
         BeginProperty Font 
            Name            =   "ÇlÇr Çoñæí©"
            Size            =   9.75
            Charset         =   128
            Weight          =   700
            Underline       =   0   'False
            Italic          =   0   'False
            Strikethrough   =   0   'False
         EndProperty
         Height          =   900
         Index           =   1
         Left            =   120
         Picture         =   "frmMain.frx":216E
         Style           =   1  '∏ﬁ◊Ã®Ø∏Ω
         TabIndex        =   2
         Top             =   1135
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "aaa"
         BeginProperty Font 
            Name            =   "ÇlÇr Çoñæí©"
            Size            =   9.75
            Charset         =   128
            Weight          =   700
            Underline       =   0   'False
            Italic          =   0   'False
            Strikethrough   =   0   'False
         EndProperty
         Height          =   900
         Index           =   0
         Left            =   120
         Picture         =   "frmMain.frx":2478
         Style           =   1  '∏ﬁ◊Ã®Ø∏Ω
         TabIndex        =   1
         Top             =   120
         Width           =   1035
      End
   End
End
Attribute VB_GlobalNameSpace = False
Attribute VB_Creatable = False
Attribute VB_PredeclaredId = True
Attribute VB_Exposed = False
Option Explicit
