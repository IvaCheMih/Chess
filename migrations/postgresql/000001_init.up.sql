CREATE TABLE users (
    id           serial       NOT NULL,
    telegram_id  integer      NOT NULL default 0,
    password     varchar(100) NOT NULL,

    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE board_cells (
    id           serial       NOT NULL,
    game_id      integer      NOT NULL default 0,
    index_cell   integer      NOT NULL default 0,
    figure_id    integer      NOT NULL default 0,

    CONSTRAINT board_cells_pkey PRIMARY KEY (id)
);

CREATE TYPE GAME_ENDGAME AS ENUM ('NotEndgame', 'Draw', 'Mate', 'Pat', 'Repetition',  'NoLosses', 'GiveUp');
CREATE TYPE GAME_STATUS AS ENUM ('Created', 'Cancelled', 'Active', 'Ended');

CREATE TABLE games (
    id                      serial        NOT NULL,
    white_user_id           integer       NOT NULL default 0,
    black_user_id           integer       NOT NULL default 0,
    status                  GAME_STATUS   NOT NULL default 'Created',

    end_reason              GAME_ENDGAME  NOT NULL default 'NotEndgame',
    winner_user_id          integer       NOT NULL default 0,

    is_check_white          BOOLEAN       NOT NULL default false,
    white_king_castling     BOOLEAN       NOT NULL default false,

    white_rook_a_castling   BOOLEAN       NOT NULL default false,
    white_rook_h_castling   BOOLEAN       NOT NULL default false,

    is_check_black          BOOLEAN       NOT NULL default false,
    black_king_castling     BOOLEAN       NOT NULL default false,

    black_rook_a_castling   BOOLEAN       NOT NULL default false,
    black_rook_h_castling   BOOLEAN       NOT NULL default false,

    last_loss               integer       NOT NULL default 0,
    last_pawn_move          integer       ,
    side                    BOOLEAN       NOT NULL default true,

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

      is_check_black        BOOLEAN      NOT NULL default false,


      CONSTRAINT moves_pkey PRIMARY KEY (id)
);


