VERSION 5.00
Object = "{00F72214-59C6-45B0-840A-F09FDEB2D5F4}#3.4#0"; "PhoneAPI.ocx"
Object = "{7571CEA4-BD56-40F8-8C6B-4E938351380F}#1.0#0"; "afValue.ocx"
Begin VB.Form frmMain 
   BackColor       =   &H00008000&
   BorderStyle     =   3  '固定ﾀﾞｲｱﾛｸﾞ
   ClientHeight    =   13455
   ClientLeft      =   45
   ClientTop       =   45
   ClientWidth     =   1470
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
         Caption         =   "システム"
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
         Index           =   12
         Left            =   120
         Picture         =   "frmMain.frx":0000
         Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
         TabIndex        =   13
         Top             =   12300
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "マスタ"
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
         Index           =   11
         Left            =   120
         Picture         =   "frmMain.frx":030A
         Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
         TabIndex        =   12
         Top             =   11285
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "ユーザ"
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
         Index           =   10
         Left            =   120
         Picture         =   "frmMain.frx":0614
         Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
         TabIndex        =   11
         Top             =   10270
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "専用業務"
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
         Index           =   9
         Left            =   120
         Picture         =   "frmMain.frx":091E
         Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
         TabIndex        =   10
         Top             =   9255
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "CTI"
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
         Index           =   8
         Left            =   120
         Picture         =   "frmMain.frx":0C28
         Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
         TabIndex        =   9
         Top             =   8240
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "OB管理"
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
         Index           =   7
         Left            =   120
         Picture         =   "frmMain.frx":0F32
         Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
         TabIndex        =   8
         Top             =   7225
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "アイテム"
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
         Index           =   6
         Left            =   120
         Picture         =   "frmMain.frx":123C
         Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
         TabIndex        =   7
         Top             =   6210
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "データ"
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
         Index           =   5
         Left            =   120
         Picture         =   "frmMain.frx":1546
         Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
         TabIndex        =   6
         Top             =   5195
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "帳票"
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
         Index           =   4
         Left            =   120
         Picture         =   "frmMain.frx":1850
         Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
         TabIndex        =   5
         Top             =   4180
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "申送管理"
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
         Index           =   3
         Left            =   120
         Picture         =   "frmMain.frx":1B5A
         Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
         TabIndex        =   4
         Top             =   3165
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "依頼書"
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
         Index           =   2
         Left            =   120
         Picture         =   "frmMain.frx":1E64
         Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
         TabIndex        =   3
         Top             =   2145
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "応対管理"
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
         Index           =   1
         Left            =   120
         Picture         =   "frmMain.frx":216E
         Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
         TabIndex        =   2
         Top             =   1135
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "メイン"
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
         Index           =   0
         Left            =   120
         Picture         =   "frmMain.frx":2478
         Style           =   1  'ｸﾞﾗﾌｨｯｸｽ
         TabIndex        =   1
         Top             =   120
         Width           =   1035
      End
   End
End
Attribute VB_Name = "frmMain"
Attribute VB_GlobalNameSpace = False
Attribute VB_Creatable = False
Attribute VB_PredeclaredId = True
Attribute VB_Exposed = False
Option Explicit

Private Sub mBtnControl()

    On Error GoTo Err_mBtnControl

    '    Public Const FUNC_000401 As String = "000401"    ' [応対管理]
    '    Public Const FUNC_000402 As String = "000402"    ' [依頼書]
    '    Public Const FUNC_000403 As String = "000403"    ' [申送管理]
    '    Public Const FUNC_000404 As String = "000404"    ' [帳票]
    '    Public Const FUNC_000405 As String = "000405"    ' [データ]
    '    Public Const FUNC_000406 As String = "000406"    ' [アイテム]
    '    Public Const FUNC_000407 As String = "000407"    ' [OB管理]
    '    Public Const FUNC_000408 As String = "000408"    ' [CTI]
    '    Public Const FUNC_000409 As String = "000409"    ' [専用業務]
    '    Public Const FUNC_000410 As String = "000410"    ' [ユーザ]
    '    Public Const FUNC_000411 As String = "000411"    ' [マスタ]
    '    Public Const FUNC_000412 As String = "000412"    ' [システム]

    If AF_FuncCheck(FUNC_000401, afValue) = False Then btnMenu(1).Enabled = False
    If AF_FuncCheck(FUNC_000402, afValue) = False Then btnMenu(2).Enabled = False
    If AF_FuncCheck(FUNC_000403, afValue) = False Then btnMenu(3).Enabled = False
    If AF_FuncCheck(FUNC_000404, afValue) = False Then btnMenu(4).Enabled = False
    If AF_FuncCheck(FUNC_000405, afValue) = False Then btnMenu(5).Enabled = False
    If AF_FuncCheck(FUNC_000406, afValue) = False Then btnMenu(6).Enabled = False
    If AF_FuncCheck(FUNC_000407, afValue) = False Then btnMenu(7).Enabled = False
    If AF_FuncCheck(FUNC_000408, afValue) = False Then btnMenu(8).Enabled = False
    If AF_FuncCheck(FUNC_000409, afValue) = False Then btnMenu(9).Enabled = False
    If AF_FuncCheck(FUNC_000410, afValue) = False Then btnMenu(10).Enabled = False
    If AF_FuncCheck(FUNC_000411, afValue) = False Then btnMenu(11).Enabled = False
    If AF_FuncCheck(FUNC_000412, afValue) = False Then btnMenu(12).Enabled = False

Exit_mBtnControl:
    Exit Sub

Err_mBtnControl:
    MsgBox Err.Description
    Resume Exit_mBtnControl

End Sub

'---------------------------------------------------------------------
' 目的説明  -- ボタン押下
' 引数      --
' 戻り値    -- なし
'---------------------------------------------------------------------
Private Sub btnMenu_Click(Index As Integer)
    On Error GoTo Err_btnMenu_Click

    Dim Rq As Long
    Dim strAppl As String
    Dim cnn As New ADODB.Connection

    strAppl = ""

    cnn.Open gDbConnectionString

    Select Case Index

    Case 0  ' メイン
        If ucPhoneApi1.MultiMain = 0 Or gLogonMode <> 9 Then
            '「メイン」単独起動モードまたはアウトバウンドのときは起動時にボタンDisableにする
            btnMenu(Index).Enabled = False
            TraceLog LOG_DBG, "「メイン」ボタン無効"
        Else
            btnMenu(Index).Enabled = True
            TraceLog LOG_DBG, "「メイン」ボタン有効"
        End If
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "000")

    Case 1
        ' 応対管理
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "010")
        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000401, True, cnn

    Case 2  ' 依頼書
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "020")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000402, True, cnn

    Case 3  ' 申送事項
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "030")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000403, True, cnn

    Case 4  ' 帳票
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "040")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000404, True, cnn

    Case 5  ' データ
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "050")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000405, True, cnn

    Case 6  ' アイテム
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "060")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000406, True, cnn

    Case 7  ' OB管理
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "070")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000407, True, cnn

    Case 8  ' CTI
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "080")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000408, True, cnn

    Case 9  ' 専用業務
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "090")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000409, True, cnn

    Case 10  ' ユーザ
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "100")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000410, True, cnn

    Case 11    ' マスタ
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "110")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000411, True, cnn

    Case 12    ' システム
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "120")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000412, True, cnn

    End Select

    cnn.Close

    If strAppl <> "" Then
        TraceLog LOG_INF, "起動アプリ: " & strAppl
        Rq = Shell(App.Path & "\" & strAppl & " " & gUser.UserID, vbNormalFocus)
    End If

Exit_btnMenu_Click:
    Exit Sub

Err_btnMenu_Click:
    ErrorBreak "btnMenu_Click"
    Resume Exit_btnMenu_Click

End Sub

'---------------------------------------------------------------------
' 目的説明  -- データ受信
' 引数      --
' 戻り値    -- なし
'---------------------------------------------------------------------
Private Sub afValue_OnData(strData As String)

    TraceLog LOG_DBG, Me.Name & ".OnData=" & strData

    Select Case strData
        '----------------------------------
    Case "LOCK"         'ロック中(作業,離席,休憩,研修)
        '----------------------------------
    Case "UNLOCK"       'ロック解除
        '----------------------------------
    Case "ENDLOCK"      'ロック認証せずに終了(起動中のアプリを終了させる)
        '終了後処理処理
        ExitProcedure
        End
        '----------------------------------
    Case "ENDPHONE"     '終了させる
        '終了後処理処理
        ExitProcedure
        End

        '----------------------------------
    Case "ENDSYSTEM"    'すべて終了させる
        Unload Me
        Exit Sub

        '----------------------------------
    Case "CHANGEGROUP"  '業務切替

        'ログオン情報を読込み
        LoadLogonData
        '再起動
        Form_Load

    End Select

End Sub

'---------------------------------------------------------------------
' 目的説明  -- 変数変更
' 引数      --
' 戻り値    -- なし
'---------------------------------------------------------------------
Private Sub afValue_OnValue(strName As String, strData As String)

    Select Case strName

    Case "ReadMessage"  '申送事項
        If strData = "1" Then
            btnMenu(3).BackColor = cRED3    '未読あり
        Else
            btnMenu(3).BackColor = cGRAY
        End If

    Case "ReadRequest"  '依頼書
        If strData = "1" Then
            btnMenu(2).BackColor = cRED3    '未読あり
        Else
            btnMenu(2).BackColor = cGRAY
        End If

    End Select

End Sub

Private Sub GetUserFunctionPower()

    On Error GoTo Err_GetUserFunctionPower

    Dim FuncIDList As uUserFuncitonPower

    Dim UserFunctionID As String
    Dim FuncCategories As String

    Dim cnn As New ADODB.Connection
    Dim rs As New ADODB.Recordset
    Dim SQL As String

    cnn.Open gDbConnectionString

    SQL = "select FunctionID from TB_USERFUNCTIONPOWER where UserID = '" & gValue.UserID & "'"
    rs.Open SQL, cnn

    Do Until rs.EOF

        UserFunctionID = rs!FunctionID
        FuncCategories = Left(UserFunctionID, 2)

        Select Case FuncCategories
        Case "00"
            FuncIDList.A00 = FuncIDList.A00 & "," & UserFunctionID
        Case "01"
            FuncIDList.A01 = FuncIDList.A01 & "," & UserFunctionID
        Case "02"
            FuncIDList.A02 = FuncIDList.A02 & "," & UserFunctionID
        Case "03"
            FuncIDList.A03 = FuncIDList.A03 & "," & UserFunctionID
        Case "04"
            FuncIDList.A04 = FuncIDList.A04 & "," & UserFunctionID
        Case "05"
            FuncIDList.A05 = FuncIDList.A05 & "," & UserFunctionID
        Case "06"
            FuncIDList.A06 = FuncIDList.A06 & "," & UserFunctionID
        Case "07"
            FuncIDList.A07 = FuncIDList.A07 & "," & UserFunctionID
        Case "08"
            FuncIDList.A08 = FuncIDList.A08 & "," & UserFunctionID
        Case "09"
            FuncIDList.A09 = FuncIDList.A09 & "," & UserFunctionID
        Case "10"
            FuncIDList.A10 = FuncIDList.A10 & "," & UserFunctionID
        Case "11"
            FuncIDList.A11 = FuncIDList.A11 & "," & UserFunctionID
        Case "12"
            FuncIDList.A12 = FuncIDList.A12 & "," & UserFunctionID
        Case "13"
            FuncIDList.A13 = FuncIDList.A13 & "," & UserFunctionID
        End Select

        rs.MoveNext

    Loop

    rs.Close
    cnn.Close

    afValue.v("A00") = Mid(FuncIDList.A00, 2)
    afValue.v("A01") = Mid(FuncIDList.A01, 2)
    afValue.v("A02") = Mid(FuncIDList.A02, 2)
    afValue.v("A03") = Mid(FuncIDList.A03, 2)
    afValue.v("A04") = Mid(FuncIDList.A04, 2)
    afValue.v("A05") = Mid(FuncIDList.A05, 2)
    afValue.v("A06") = Mid(FuncIDList.A06, 2)
    afValue.v("A07") = Mid(FuncIDList.A07, 2)
    afValue.v("A08") = Mid(FuncIDList.A08, 2)
    afValue.v("A09") = Mid(FuncIDList.A09, 2)
    afValue.v("A10") = Mid(FuncIDList.A10, 2)
    afValue.v("A11") = Mid(FuncIDList.A11, 2)
    afValue.v("A12") = Mid(FuncIDList.A12, 2)
    afValue.v("A13") = Mid(FuncIDList.A13, 2)

Exit_GetUserFunctionPower:
    Exit Sub

Err_GetUserFunctionPower:
    MsgBox Err.Description
    Resume Exit_GetUserFunctionPower

End Sub

'---------------------------------------------------------------------
' 目的説明  -- 「frmMain」フォームロード
' 引数      --
' 戻り値    -- なし
'---------------------------------------------------------------------
Private Sub Form_Load()

    On Error GoTo Err_Form_Load

    Dim Rq As Long
    Dim strCom As String
    Dim strAppl As String

    Me.Move 17650, 0

    '電話API接続
    ucPhoneApi1.OpenTel "False"

    'ユーザーグローバル情報セット (業務開始時に構造体の値をAFグローバル変数に代入します。)
    gValue.LogonMode = gLogonMode

    'AFグローバル変数に代入
    afValue.v("UserID") = gValue.UserID
    afValue.v("OpeName") = gValue.UserName
    afValue.v("PCName") = gValue.PCName
    afValue.v("Extension") = gValue.Extension
    afValue.v("LogonMode") = gValue.LogonMode
    afValue.v("ClientCode") = gValue.ClientCode
    afValue.v("ClientName") = gValue.ClientName
    afValue.v("GroupCode") = gValue.GroupCode
    afValue.v("JobCode") = gValue.JobCode
    afValue.v("ListCode") = gValue.ListCode
    'START 20210525 新リストコール
    afValue.v("LogonSubMode") = gValue.LogonSubMode
    afValue.v("MainInOut") = gValue.MainInOut
    'END 20210525 新リストコール

    TraceLog LOG_DBG, "UserID[" & gValue.UserID & "] UerName[" & gValue.UserName & "] PC[" & gValue.PCName & "] Ext[" & gValue.Extension & "] Logon[" & gValue.LogonMode & "] ClientCode[" & gValue.ClientCode & "] ClientName[" & gValue.ClientName & "] Group[" & gValue.GroupCode & "] Job[" & gValue.JobCode & "] Lsit[" & gValue.ListCode & "]"

    ' 機能権限をAFｸﾞﾛｰﾊﾞﾙ変数に代入
    GetUserFunctionPower

    'ログレベル(0=なし, 1=ERR, 2=ERR+WRN, 3=ERR+WRN+INF, 4=ERR+WRN+INF+SQL, 5以上=ALL)
    afValue.v("LogLevel") = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_APPL, "LogLevel")
    If afValue.v("LogLevel") = "" Then
        afValue.v("LogLevel") = "9"
    End If
    gValue.LogLevel = CInt(afValue.v("LogLevel"))


    'SV電話機を起動する (UserID,作業状態指定,電話業務種別)
    '   作業状態指定 (0=ログオフ(作業中),1=無作業,2=待機(着信待ち),3=作業)
    '   電話業務種別 (0=未ログイン,1=PD,2=AC,3=PC,4=DPC,5=PV,6=DPV,7:=DC,8=DDC,9=IB)

    strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "200")
    If strAppl = "" Then strAppl = "AF_SVPHONE.exe"
    If strAppl = "AF_PHONECTL_NON.exe" Then strAppl = "AF_SVPHONE.exe"  '2009/5/27変更したので前値と互換維持動作
    strCom = gUser.UserID & "," & gWorkStatus & "," & gLogonMode
    Rq = Shell(App.Path & "\" & strAppl & " " & strCom, vbNormalFocus)

    ' ﾎﾞﾀﾝ制御
    mBtnControl

    'メニュー表示
    Me.Show
    DoEvents

    'アウトバウンドのときは「メイン」を起動する
    ' (0=未ログイン,1=PD,2=AC,3=PC,4=DPC,5=PV,6=DPV,7:=DC,8=DDC,9=IB,10=TEL)
    '電話機単体以外では「メイン」を起動する
    If gValue.LogonMode <> 10 Then
        '複数メイン対応 (0=１つのみ、1=複数可能) - 電話機の起動が間に合わず値が正しく入らないのでここで取得
        ucPhoneApi1.MultiMain = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_APPL, "MultiMain")
        '「メイン」ボタン押下
        Call btnMenu_Click(0)
    Else
        '電話機単体起動
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "270")
        If strAppl = "" Then strAppl = "AF_PHONE.exe"
        Rq = Shell(App.Path & "\" & strAppl & " " & gUser.UserID, vbNormalFocus)
    End If


Exit_Form_Load:
    Exit Sub

Err_Form_Load:
    ErrorBreak "Load"
    Resume Exit_Form_Load

End Sub

'---------------------------------------------------------------------
' 目的説明  -- 「frmMain」フォームアンロード
' 引数      --
' 戻り値    -- なし
'---------------------------------------------------------------------
Private Sub Form_Unload(Cancel As Integer)

    On Error GoTo Err_Form_Unload

    '電話API切断
    ucPhoneApi1.CloseTel

    '終了後処理処理
    ExitProcedure

    End

Exit_Form_Unload:
    Exit Sub

Err_Form_Unload:
    ErrorBreak "Unload"
    Resume Exit_Form_Unload

End Sub

'---------------------------------------------------------------------
' 目的説明  -- 電話変数変更
' 引数      --
' 戻り値    -- なし
'---------------------------------------------------------------------
Private Sub ucPhoneApi1_OnValue(strName As String, strData As String)

    Select Case strName

    Case "MainIndex"                '電話機と「メイン」の結び付け変更
        '「メイン」は単独使用モード？
        If ucPhoneApi1.MultiMain = 0 Or gLogonMode <> 9 Then
            If ucPhoneApi1.MainIndex = "" Or ucPhoneApi1.MainIndex = "0" Then
                btnMenu(0).Enabled = True       '未稼働であれば起動可能
                TraceLog LOG_DBG, "「メイン」ボタン有効"
            Else
                btnMenu(0).Enabled = False      '電話機と結びつきがあれば起動不可
                TraceLog LOG_DBG, "「メイン」ボタン無効"
            End If
        End If

    End Select

End Sub



'---------------------------------------------------------------------
' 目的説明  -- ログインに必要なデータをAFValueから読込む
' 引数      --
' 戻り値    -- なし
'---------------------------------------------------------------------
Private Sub LoadLogonData()

'AFグローバル変数に代入
    gValue.UserID = afValue.v("UserID")
    gValue.UserName = afValue.v("OpeName")
    gValue.PCName = afValue.v("PCName")
    gValue.Extension = afValue.v("Extension")

    gLogonMode = afValue.v("LogonMode")  'ログオンモードはグローバルへ
    'START 20210525 新リストコール
    gValue.LogonSubMode = afValue.v("LogonSubMode")
    gValue.MainInOut = afValue.v("MainInOut")
    'END 20210525 新リストコール

    gValue.ClientCode = afValue.v("ClientCode")
    gValue.ClientName = afValue.v("ClientName")
    gValue.GroupCode = afValue.v("GroupCode")
    gValue.JobCode = afValue.v("JobCode")
    gValue.ListCode = afValue.v("ListCode")


    TraceLog LOG_DBG, "UserID[" & gValue.UserID & "] UerName[" & gValue.UserName & "] PC[" & gValue.PCName & "] Ext[" & gValue.Extension & "] Logon[" & gLogonMode & "] ClientCode[" & gValue.ClientCode & "] ClientName[" & gValue.ClientName & "] Group[" & gValue.GroupCode & "] Job[" & gValue.JobCode & "] Lsit[" & gValue.ListCode & "]"

End Sub
