CREATE TABLE users (
    id         serial       NOT NULL,
    password   varchar(100) NOT NULL,

    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE boardCells (
    id          serial       NOT NULL,
    gameId      integer      NOT NULL default 0,
    indexCell   integer      NOT NULL default 0,
    figureId    integer      NOT NULL default 0,

    CONSTRAINT boardCells_pkey PRIMARY KEY (id)
);

CREATE TABLE figures (
   id          serial        NOT NULL,
   type        varchar(100)  NOT NULL,
   isWhite     BOOLEAN       NOT NULL,

   CONSTRAINT figures_pkey PRIMARY KEY (id)
);

CREATE TABLE games (
    id              serial       NOT NULL,
    whiteUserId     integer      NOT NULL default 0,
    blackUserId     integer      NOT NULL default 0,
    isStarted       BOOLEAN      NOT NULL default false,
    isEnded         BOOLEAN      NOT NULL default false,
    isCheckWhite    BOOLEAN      NOT NULL default false,
    whiteKingCell   integer      NOT NULL default 60,
    isCheckBlack    BOOLEAN      NOT NULL default false,
    blackKingCell   integer      NOT NULL default 4,
    side            integer      NOT NULL default 1,

    CONSTRAINT games_pkey PRIMARY KEY (id)
);

CREATE TABLE moves (
      id               serial       NOT NULL,
      gameId           integer      NOT NULL default 0,
      moveNumber       integer      NOT NULL default 0,
      from_id          integer      NOT NULL default 0,
      to_id            integer      NOT NULL default 0,
      figureId         integer      NOT NULL default 0,
      killedFigureId   integer      NOT NULL default 0,
      newFigureId      integer      NOT NULL default 0,
      isCheckWhite     BOOLEAN      NOT NULL default false,
      whiteKingCell    integer      NOT NULL default 0,
      isCheckBlack     BOOLEAN      NOT NULL default false,
      blackKingCell    integer      NOT NULL default 0,

      CONSTRAINT moves_pkey PRIMARY KEY (id)
);


INSERT INTO figures (id, type, isWhite) VALUES (1, 'r', false);
INSERT INTO figures (id, type, isWhite) VALUES (2, 'k', false);
INSERT INTO figures (id, type, isWhite) VALUES (3, 'b', false);
INSERT INTO figures (id, type, isWhite) VALUES (4, 'q', false);
INSERT INTO figures (id, type, isWhite) VALUES (5, 'K', false);
INSERT INTO figures (id, type, isWhite) VALUES (6, 'p', false);

INSERT INTO figures (id, type, isWhite) VALUES (7, 'r', true);
INSERT INTO figures (id, type, isWhite) VALUES (8, 'k', true);
INSERT INTO figures (id, type, isWhite) VALUES (9, 'b', true);
INSERT INTO figures (id, type, isWhite) VALUES (10, 'q', true);
INSERT INTO figures (id, type, isWhite) VALUES (11, 'K', true);
INSERT INTO figures (id, type, isWhite) VALUES (12, 'p', true);

-- insert into moves (gameId,moveNumber, from_id,to_id,figureId, killedFigureId, newFigureId, isCheckWhite , whiteKingCell, isCheckBlack, blackKingCell)
--         values (0, 0, 0, 0, 0, 0, 0, false,0,false,0);

