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
    id          serial        NOT NULL,
    whiteUserId integer       NOT NULL default 0,
    blackUserId integer       NOT NULL default 0,
    isStarted   BOOLEAN       NOT NULL default false,
    isEnded     BOOLEAN       NOT NULL default false,

    CONSTRAINT games_pkey PRIMARY KEY (id)
);

CREATE TABLE moves (
      id               serial       NOT NULL,
      gameId           integer      NOT NULL,
      moveNumber       integer      NOT NULL,
      from_id          integer      NOT NULL,
      to_id            integer      NOT NULL,
      figureId         integer      NOT NULL,
      killedFigureId   integer      NOT NULL,
      newFigureId      integer      NOT NULL,
      isCheckWhite     BOOLEAN      NOT NULL,
      whiteKingCell     integer      NOT NULL,
      isCheckBlack     BOOLEAN      NOT NULL,
      blackKingCell     integer      NOT NULL,

      CONSTRAINT moves_pkey PRIMARY KEY (id)
);


INSERT INTO figures (id, type, isWhite) VALUES (1, 'r', false);
INSERT INTO figures (id, type, isWhite) VALUES (2, 'h', false);
INSERT INTO figures (id, type, isWhite) VALUES (3, 'b', false);
INSERT INTO figures (id, type, isWhite) VALUES (4, 'q', false);
INSERT INTO figures (id, type, isWhite) VALUES (5, 'k', false);
INSERT INTO figures (id, type, isWhite) VALUES (6, 'p', false);

INSERT INTO figures (id, type, isWhite) VALUES (7, 'r', true);
INSERT INTO figures (id, type, isWhite) VALUES (8, 'h', true);
INSERT INTO figures (id, type, isWhite) VALUES (9, 'b', true);
INSERT INTO figures (id, type, isWhite) VALUES (10, 'q', true);
INSERT INTO figures (id, type, isWhite) VALUES (11, 'k', true);
INSERT INTO figures (id, type, isWhite) VALUES (12, 'p', true);
