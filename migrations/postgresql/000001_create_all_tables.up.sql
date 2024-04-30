CREATE TABLE users (
    id         serial       NOT NULL,
    password   varchar(100) NOT NULL,

    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE board_cells (
    id           serial       NOT NULL,
    game_id      integer      NOT NULL default 0,
    index_cell   integer      NOT NULL default 0,
    figure_id    integer      NOT NULL default 0,

    CONSTRAINT board_cells_pkey PRIMARY KEY (id)
);

CREATE TABLE figures (
   id           serial        NOT NULL,
   type         varchar(100)  NOT NULL,
   is_white     BOOLEAN       NOT NULL,

   CONSTRAINT figures_pkey PRIMARY KEY (id)
);

CREATE TABLE games (
    id                      serial       NOT NULL,
    white_user_id           integer      NOT NULL default 0,
    black_user_id           integer      NOT NULL default 0,
    is_started              BOOLEAN      NOT NULL default false,
    is_ended                BOOLEAN      NOT NULL default false,

    is_check_white          BOOLEAN      NOT NULL default false,
    white_king_cell         integer      NOT NULL default 60,
    white_king_castling     BOOLEAN      NOT NULL default false,

    white_rook_a_castling   BOOLEAN      NOT NULL default false,
    white_rook_h_castling   BOOLEAN      NOT NULL default false,

    is_check_black          BOOLEAN      NOT NULL default false,
    black_king_cell         integer      NOT NULL default 4,
    black_king_castling     BOOLEAN      NOT NULL default false,

    black_rook_a_castling   BOOLEAN      NOT NULL default false,
    black_rook_h_castling   BOOLEAN      NOT NULL default false,

    side                    integer      NOT NULL default 1,

    CONSTRAINT games_pkey PRIMARY KEY (id)
);

CREATE TABLE moves (
      id                    serial       NOT NULL,
      game_id               integer      NOT NULL default 0,
      move_number           integer      NOT NULL default 0,
      from_id               integer      NOT NULL default 0,
      to_id                 integer      NOT NULL default 0,
      figure_id             integer      NOT NULL default 0,
      killed_figure_id      integer      NOT NULL default 0,
      new_figure_id         integer      NOT NULL default 0,

      is_check_white        BOOLEAN      NOT NULL default false,
      white_king_cell       integer      NOT NULL default 0,

      is_check_black        BOOLEAN      NOT NULL default false,
      black_king_cell       integer      NOT NULL default 0,


      CONSTRAINT moves_pkey PRIMARY KEY (id)
);


INSERT INTO figures (id, type, is_white) VALUES (1, 'r', false);
INSERT INTO figures (id, type, is_white) VALUES (2, 'k', false);
INSERT INTO figures (id, type, is_white) VALUES (3, 'b', false);
INSERT INTO figures (id, type, is_white) VALUES (4, 'q', false);
INSERT INTO figures (id, type, is_white) VALUES (5, 'K', false);
INSERT INTO figures (id, type, is_white) VALUES (6, 'p', false);

INSERT INTO figures (id, type, is_white) VALUES (7, 'r', true);
INSERT INTO figures (id, type, is_white) VALUES (8, 'k', true);
INSERT INTO figures (id, type, is_white) VALUES (9, 'b', true);
INSERT INTO figures (id, type, is_white) VALUES (10, 'q', true);
INSERT INTO figures (id, type, is_white) VALUES (11, 'K', true);
INSERT INTO figures (id, type, is_white) VALUES (12, 'p', true);

