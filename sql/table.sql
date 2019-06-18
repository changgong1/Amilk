use amilk;
DROP table if exists MovieCardTbl;
Create table MovieCardTbl (
    ID  INT primary key auto_increment,
    SourceName       VARCHAR(64)  COMMENT 'movieName',
    TransNameID      INT          COMMENT '其他名字',
    ScoreType        INT          COMMENT '得分类别',             # 1-好, 2-中，3-坏
    Score            VARCHAR(4)   COMMENT '得分',
    ScreenDate       DATETIME     COMMENT '上映日期',
    DirectIDs        VARCHAR(8)   COMMENT '导演',
    ActorIDs         VARCHAR(16)  COMMENT '演员',
    ProducerIDs      VARCHAR(8)   COMMENT '制片人',
    PlaywrightIDs    VARCHAR(8)   COMMENT '编剧',
    TagSourceName    VARCHAR(20)  COMMENT '电影类别',
    TagTransNameIDs  VARCHAR(20)  COMMENT '电影类别',
    Outline          TEXT         COMMENT '故事梗概',
    HonourIDs        VARCHAR(8)   COMMENT '荣誉',
    ImageIDs         VARCHAR(256) COMMENT '图片地址',
    GbScore          FLOAT        COMMENT '高分',
    MbScore          FLOAT        COMMENT '中分',
    BdScore          FLOAT        COMMENT '低分'
)
auto_increment = 1
ENGINE=INNODB default  CHARSET=utf8 COMMENT '电影名片表';

DROP table if exists ArtistCardTbl;
Create TABLE ArtistCardTbl (
    ID INT primary key auto_increment,
    UserID          INT          COMMENT '用户ID',            # 如果该艺人是Amilk用户，则有该内容
    SourceName      VARCHAR(64)  COMMENT '名字',
    TransNameID     INT          COMMENT '其他名字',
    Birthdate       DATETIME     COMMENT '出生日期',
    Country         VARCHAR(8)   COMMENT '国家',
    CountryID       INT          COMMENT '国家多语言',
    Positions       VARCHAR(64)  COMMENT '身份(演员,导演)',
    MasterpieceIDs  VARCHAR(8)   COMMENT '代表作',
    Introduce       TEXT         COMMENT '艺人简介',
    HonourIDs       VARCHAR(8)   COMMENT '荣誉',
    ImageIDs        VARCHAR(256) COMMENT '图片地址'
)
auto_increment = 1
ENGINE=INNODB default CHARSET=utf8 COMMENT '艺人名片表';

DROP table if exists ManyLanguageTbl;       # 该表同时是名词表，专有名词根据ID在这里查找对应语言的释译
Create TABLE ManyLanguageTbl (
    ID INT primary key auto_increment,
    SourceTitle   VARCHAR(64) COMMENT '名字',
    ChineseTitle  VARCHAR(64) COMMENT '中文',
    EnglishTitle  VARCHAR(64) COMMENT '英文',
    ArabiaTitle   VARCHAR(64) COMMENT '阿拉伯文',
    JapanTitle    VARCHAR(64) COMMENT '日文',
    KoreanTitle   VARCHAR(64) COMMENT '韩文',
    ThailandTitle VARCHAR(64) COMMENT '泰语'
)
auto_increment = 1
ENGINE=INNODB default CHARSET=utf8 COMMENT '多语言表';

DROP table if exists HonourCardTbl;
Create table HonourCardTbl(
    ID INT primary key auto_increment,
    StructureSourceName     VARCHAR(256) COMMENT '机构名称',
    StructureTransNameID    INT          COMMENT '机构多语言',
    StructureSourceTitle    VARCHAR(256) COMMENT '对外官称',
    StructureTransTitleID   INT          COMMENT '对外官称多语言',
    UnitSourceName          VARCHAR(256) COMMENT '单元名称',
    UnitTransNameID         INT          COMMENT '单元名称多语言',
    GlorySourceName         VARCHAR(256) COMMENT '荣誉称号',
    GloryTransNameID        INT          COMMENT '荣誉称号多语言',
    MovieCardIDs            VARCHAR(8)   COMMENT '获奖作品ID',
    ArtistCardID            VARCHAR(8)   COMMENT '获奖人物对象'
)   
auto_increment = 1
ENGINE=INNODB default CHARSET=utf8 COMMENT '荣誉表';

DROP table if exists UsersTbl;
Create table UsersTbl(
    ID INT primary key auto_increment,
    UserID             VARCHAR(8)   COMMENT '登陆ID',
    Password           VARCHAR(32)  COMMENT '密码',
    UserSourceName     VARCHAR(256) COMMENT '用户名',
    UserTransNameID    INT          COMMENT '名称多语言',
    Birthdate          DATETIME     COMMENT '出生日期',
    CreateTime         DATETIME     COMMENT '创建时间'
)
auto_increment = 1
ENGINE=INNODB default CHARSET=utf8 COMMENT '用户表';

DROP table if exists UserCriterionTbl;
Create table UserCriterionTbl(
    ID INT primary key auto_increment,
    UserID      INT          COMMENT '评价人',
    ObjectID    INT          COMMENT '评价对象',
    ObjectType  INT          COMMENT '对象类型',    # 0-评论回复 1-影片, 2-艺人, 
    ScoreType   INT          COMMENT '评分类型',
    Comment     TEXT         COMMENT '评价内容',
    Article     TEXT         COMMENT '长文评价',
    ImageIDs    VARCHAR(256) COMMENT '图片地址',
    CreateTime  DATETIME     COMMENT '创建时间'
)
auto_increment = 1
ENGINE=INNODB default CHARSET=utf8 COMMENT '动态表';

DROP table if exists ImagesTbl;
Create table ImagesTbl(
    ID INT primary key auto_increment,
    UserID            INT            COMMENT '图片上传者',
    ImageDir          VARCHAR(256)   COMMENT '图片地址',
    ImageSourceName   VARCHAR(256)   COMMENT '图片名称',
    ImageTransNameID  VARCHAR(8)     COMMENT '图片名称多语言',
    TagSourceName     VARCHAR(256)   COMMENT '图片标签',
    TagTransNameID    INT            COMMENT '标签多语言'
)
auto_increment = 1
ENGINE=INNODB default CHARSET=utf8 COMMENT '动态表';