VERSION 5.00
Object = "{00F72214-59C6-45B0-840A-F09FDEB2D5F4}#3.4#0"; "PhoneAPI.ocx"
Object = "{7571CEA4-BD56-40F8-8C6B-4E938351380F}#1.0#0"; "afValue.ocx"
Begin VB.Form frmMain 
   BackColor       =   &H00008000&
   BorderStyle     =   3  '�Œ��޲�۸�
   ClientHeight    =   13455
   ClientLeft      =   45
   ClientTop       =   45
   ClientWidth     =   1470
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
         Caption         =   "�V�X�e��"
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
         Index           =   12
         Left            =   120
         Picture         =   "frmMain.frx":0000
         Style           =   1  '���̨���
         TabIndex        =   13
         Top             =   12300
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "�}�X�^"
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
         Index           =   11
         Left            =   120
         Picture         =   "frmMain.frx":030A
         Style           =   1  '���̨���
         TabIndex        =   12
         Top             =   11285
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "���[�U"
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
         Index           =   10
         Left            =   120
         Picture         =   "frmMain.frx":0614
         Style           =   1  '���̨���
         TabIndex        =   11
         Top             =   10270
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "��p�Ɩ�"
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
         Index           =   9
         Left            =   120
         Picture         =   "frmMain.frx":091E
         Style           =   1  '���̨���
         TabIndex        =   10
         Top             =   9255
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "CTI"
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
         Index           =   8
         Left            =   120
         Picture         =   "frmMain.frx":0C28
         Style           =   1  '���̨���
         TabIndex        =   9
         Top             =   8240
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "OB�Ǘ�"
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
         Index           =   7
         Left            =   120
         Picture         =   "frmMain.frx":0F32
         Style           =   1  '���̨���
         TabIndex        =   8
         Top             =   7225
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "�A�C�e��"
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
         Index           =   6
         Left            =   120
         Picture         =   "frmMain.frx":123C
         Style           =   1  '���̨���
         TabIndex        =   7
         Top             =   6210
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "�f�[�^"
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
         Index           =   5
         Left            =   120
         Picture         =   "frmMain.frx":1546
         Style           =   1  '���̨���
         TabIndex        =   6
         Top             =   5195
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "���["
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
         Index           =   4
         Left            =   120
         Picture         =   "frmMain.frx":1850
         Style           =   1  '���̨���
         TabIndex        =   5
         Top             =   4180
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "�\���Ǘ�"
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
         Index           =   3
         Left            =   120
         Picture         =   "frmMain.frx":1B5A
         Style           =   1  '���̨���
         TabIndex        =   4
         Top             =   3165
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "�˗���"
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
         Index           =   2
         Left            =   120
         Picture         =   "frmMain.frx":1E64
         Style           =   1  '���̨���
         TabIndex        =   3
         Top             =   2145
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "���ΊǗ�"
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
         Index           =   1
         Left            =   120
         Picture         =   "frmMain.frx":216E
         Style           =   1  '���̨���
         TabIndex        =   2
         Top             =   1135
         Width           =   1035
      End
      Begin VB.CommandButton btnMenu 
         Caption         =   "���C��"
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
         Index           =   0
         Left            =   120
         Picture         =   "frmMain.frx":2478
         Style           =   1  '���̨���
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

    '    Public Const FUNC_000401 As String = "000401"    ' [���ΊǗ�]
    '    Public Const FUNC_000402 As String = "000402"    ' [�˗���]
    '    Public Const FUNC_000403 As String = "000403"    ' [�\���Ǘ�]
    '    Public Const FUNC_000404 As String = "000404"    ' [���[]
    '    Public Const FUNC_000405 As String = "000405"    ' [�f�[�^]
    '    Public Const FUNC_000406 As String = "000406"    ' [�A�C�e��]
    '    Public Const FUNC_000407 As String = "000407"    ' [OB�Ǘ�]
    '    Public Const FUNC_000408 As String = "000408"    ' [CTI]
    '    Public Const FUNC_000409 As String = "000409"    ' [��p�Ɩ�]
    '    Public Const FUNC_000410 As String = "000410"    ' [���[�U]
    '    Public Const FUNC_000411 As String = "000411"    ' [�}�X�^]
    '    Public Const FUNC_000412 As String = "000412"    ' [�V�X�e��]

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
' �ړI����  -- �{�^������
' ����      --
' �߂�l    -- �Ȃ�
'---------------------------------------------------------------------
Private Sub btnMenu_Click(Index As Integer)
    On Error GoTo Err_btnMenu_Click

    Dim Rq As Long
    Dim strAppl As String
    Dim cnn As New ADODB.Connection

    strAppl = ""

    cnn.Open gDbConnectionString

    Select Case Index

    Case 0  ' ���C��
        If ucPhoneApi1.MultiMain = 0 Or gLogonMode <> 9 Then
            '�u���C���v�P�ƋN�����[�h�܂��̓A�E�g�o�E���h�̂Ƃ��͋N�����Ƀ{�^��Disable�ɂ���
            btnMenu(Index).Enabled = False
            TraceLog LOG_DBG, "�u���C���v�{�^������"
        Else
            btnMenu(Index).Enabled = True
            TraceLog LOG_DBG, "�u���C���v�{�^���L��"
        End If
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "000")

    Case 1
        ' ���ΊǗ�
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "010")
        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000401, True, cnn

    Case 2  ' �˗���
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "020")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000402, True, cnn

    Case 3  ' �\������
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "030")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000403, True, cnn

    Case 4  ' ���[
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "040")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000404, True, cnn

    Case 5  ' �f�[�^
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "050")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000405, True, cnn

    Case 6  ' �A�C�e��
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "060")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000406, True, cnn

    Case 7  ' OB�Ǘ�
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "070")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000407, True, cnn

    Case 8  ' CTI
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "080")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000408, True, cnn

    Case 9  ' ��p�Ɩ�
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "090")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000409, True, cnn

    Case 10  ' ���[�U
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "100")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000410, True, cnn

    Case 11    ' �}�X�^
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "110")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000411, True, cnn

    Case 12    ' �V�X�e��
        strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "120")

        ' AccessLog -----------
        AF_InsUserAccessLog FUNC_000412, True, cnn

    End Select

    cnn.Close

    If strAppl <> "" Then
        TraceLog LOG_INF, "�N���A�v��: " & strAppl
        Rq = Shell(App.Path & "\" & strAppl & " " & gUser.UserID, vbNormalFocus)
    End If

Exit_btnMenu_Click:
    Exit Sub

Err_btnMenu_Click:
    ErrorBreak "btnMenu_Click"
    Resume Exit_btnMenu_Click

End Sub

'---------------------------------------------------------------------
' �ړI����  -- �f�[�^��M
' ����      --
' �߂�l    -- �Ȃ�
'---------------------------------------------------------------------
Private Sub afValue_OnData(strData As String)

    TraceLog LOG_DBG, Me.Name & ".OnData=" & strData

    Select Case strData
        '----------------------------------
    Case "LOCK"         '���b�N��(���,����,�x�e,���C)
        '----------------------------------
    Case "UNLOCK"       '���b�N����
        '----------------------------------
    Case "ENDLOCK"      '���b�N�F�؂����ɏI��(�N�����̃A�v�����I��������)
        '�I���㏈������
        ExitProcedure
        End
        '----------------------------------
    Case "ENDPHONE"     '�I��������
        '�I���㏈������
        ExitProcedure
        End

        '----------------------------------
    Case "ENDSYSTEM"    '���ׂďI��������
        Unload Me
        Exit Sub

        '----------------------------------
    Case "CHANGEGROUP"  '�Ɩ��ؑ�

        '���O�I������Ǎ���
        LoadLogonData
        '�ċN��
        Form_Load

    End Select

End Sub

'---------------------------------------------------------------------
' �ړI����  -- �ϐ��ύX
' ����      --
' �߂�l    -- �Ȃ�
'---------------------------------------------------------------------
Private Sub afValue_OnValue(strName As String, strData As String)

    Select Case strName

    Case "ReadMessage"  '�\������
        If strData = "1" Then
            btnMenu(3).BackColor = cRED3    '���ǂ���
        Else
            btnMenu(3).BackColor = cGRAY
        End If

    Case "ReadRequest"  '�˗���
        If strData = "1" Then
            btnMenu(2).BackColor = cRED3    '���ǂ���
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
' �ړI����  -- �ufrmMain�v�t�H�[�����[�h
' ����      --
' �߂�l    -- �Ȃ�
'---------------------------------------------------------------------
Private Sub Form_Load()

    On Error GoTo Err_Form_Load

    Dim Rq As Long
    Dim strCom As String
    Dim strAppl As String

    Me.Move 17650, 0

    '�d�bAPI�ڑ�
    ucPhoneApi1.OpenTel "False"

    '���[�U�[�O���[�o�����Z�b�g (�Ɩ��J�n���ɍ\���̂̒l��AF�O���[�o���ϐ��ɑ�����܂��B)
    gValue.LogonMode = gLogonMode

    'AF�O���[�o���ϐ��ɑ��
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
    'START 20210525 �V���X�g�R�[��
    afValue.v("LogonSubMode") = gValue.LogonSubMode
    afValue.v("MainInOut") = gValue.MainInOut
    'END 20210525 �V���X�g�R�[��

    TraceLog LOG_DBG, "UserID[" & gValue.UserID & "] UerName[" & gValue.UserName & "] PC[" & gValue.PCName & "] Ext[" & gValue.Extension & "] Logon[" & gValue.LogonMode & "] ClientCode[" & gValue.ClientCode & "] ClientName[" & gValue.ClientName & "] Group[" & gValue.GroupCode & "] Job[" & gValue.JobCode & "] Lsit[" & gValue.ListCode & "]"

    ' �@�\������AF��۰��ٕϐ��ɑ��
    GetUserFunctionPower

    '���O���x��(0=�Ȃ�, 1=ERR, 2=ERR+WRN, 3=ERR+WRN+INF, 4=ERR+WRN+INF+SQL, 5�ȏ�=ALL)
    afValue.v("LogLevel") = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_APPL, "LogLevel")
    If afValue.v("LogLevel") = "" Then
        afValue.v("LogLevel") = "9"
    End If
    gValue.LogLevel = CInt(afValue.v("LogLevel"))


    'SV�d�b�@���N������ (UserID,��Ə�Ԏw��,�d�b�Ɩ����)
    '   ��Ə�Ԏw�� (0=���O�I�t(��ƒ�),1=�����,2=�ҋ@(���M�҂�),3=���)
    '   �d�b�Ɩ���� (0=�����O�C��,1=PD,2=AC,3=PC,4=DPC,5=PV,6=DPV,7:=DC,8=DDC,9=IB)

    strAppl = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_BUTTON, "200")
    If strAppl = "" Then strAppl = "AF_SVPHONE.exe"
    If strAppl = "AF_PHONECTL_NON.exe" Then strAppl = "AF_SVPHONE.exe"  '2009/5/27�ύX�����̂őO�l�ƌ݊��ێ�����
    strCom = gUser.UserID & "," & gWorkStatus & "," & gLogonMode
    Rq = Shell(App.Path & "\" & strAppl & " " & strCom, vbNormalFocus)

    ' ���ݐ���
    mBtnControl

    '���j���[�\��
    Me.Show
    DoEvents

    '�A�E�g�o�E���h�̂Ƃ��́u���C���v���N������
    ' (0=�����O�C��,1=PD,2=AC,3=PC,4=DPC,5=PV,6=DPV,7:=DC,8=DDC,9=IB,10=TEL)
    '�d�b�@�P�̈ȊO�ł́u���C���v���N������
    If gValue.LogonMode <> 10 Then
        '�������C���Ή� (0=�P�̂݁A1=�����\) - �d�b�@�̋N�����Ԃɍ��킸�l������������Ȃ��̂ł����Ŏ擾
        ucPhoneApi1.MultiMain = SQL_GET_PROFILE_STRING(gDbConnectionString, SECTION_APPL, "MultiMain")
        '�u���C���v�{�^������
        Call btnMenu_Click(0)
    Else
        '�d�b�@�P�̋N��
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
' �ړI����  -- �ufrmMain�v�t�H�[���A�����[�h
' ����      --
' �߂�l    -- �Ȃ�
'---------------------------------------------------------------------
Private Sub Form_Unload(Cancel As Integer)

    On Error GoTo Err_Form_Unload

    '�d�bAPI�ؒf
    ucPhoneApi1.CloseTel

    '�I���㏈������
    ExitProcedure

    End

Exit_Form_Unload:
    Exit Sub

Err_Form_Unload:
    ErrorBreak "Unload"
    Resume Exit_Form_Unload

End Sub

'---------------------------------------------------------------------
' �ړI����  -- �d�b�ϐ��ύX
' ����      --
' �߂�l    -- �Ȃ�
'---------------------------------------------------------------------
Private Sub ucPhoneApi1_OnValue(strName As String, strData As String)

    Select Case strName

    Case "MainIndex"                '�d�b�@�Ɓu���C���v�̌��ѕt���ύX
        '�u���C���v�͒P�Ǝg�p���[�h�H
        If ucPhoneApi1.MultiMain = 0 Or gLogonMode <> 9 Then
            If ucPhoneApi1.MainIndex = "" Or ucPhoneApi1.MainIndex = "0" Then
                btnMenu(0).Enabled = True       '���ғ��ł���΋N���\
                TraceLog LOG_DBG, "�u���C���v�{�^���L��"
            Else
                btnMenu(0).Enabled = False      '�d�b�@�ƌ��т�������΋N���s��
                TraceLog LOG_DBG, "�u���C���v�{�^������"
            End If
        End If

    End Select

End Sub



'---------------------------------------------------------------------
' �ړI����  -- ���O�C���ɕK�v�ȃf�[�^��AFValue����Ǎ���
' ����      --
' �߂�l    -- �Ȃ�
'---------------------------------------------------------------------
Private Sub LoadLogonData()

'AF�O���[�o���ϐ��ɑ��
    gValue.UserID = afValue.v("UserID")
    gValue.UserName = afValue.v("OpeName")
    gValue.PCName = afValue.v("PCName")
    gValue.Extension = afValue.v("Extension")

    gLogonMode = afValue.v("LogonMode")  '���O�I�����[�h�̓O���[�o����
    'START 20210525 �V���X�g�R�[��
    gValue.LogonSubMode = afValue.v("LogonSubMode")
    gValue.MainInOut = afValue.v("MainInOut")
    'END 20210525 �V���X�g�R�[��

    gValue.ClientCode = afValue.v("ClientCode")
    gValue.ClientName = afValue.v("ClientName")
    gValue.GroupCode = afValue.v("GroupCode")
    gValue.JobCode = afValue.v("JobCode")
    gValue.ListCode = afValue.v("ListCode")


    TraceLog LOG_DBG, "UserID[" & gValue.UserID & "] UerName[" & gValue.UserName & "] PC[" & gValue.PCName & "] Ext[" & gValue.Extension & "] Logon[" & gLogonMode & "] ClientCode[" & gValue.ClientCode & "] ClientName[" & gValue.ClientName & "] Group[" & gValue.GroupCode & "] Job[" & gValue.JobCode & "] Lsit[" & gValue.ListCode & "]"

End Sub
