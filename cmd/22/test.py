import collections
import timeit


class Player:

    def __init__(self, deck):
        self._deck = collections.deque(deck)
        self.win_state = False

    @property
    def deck(self):
        return tuple(self._deck)

    @property
    def cards_remaining(self):
        return len(self._deck)

    @property
    def score(self):
        if not self._deck:
            return 0
        score = 0
        for raw_multiplier, card in enumerate(reversed(self._deck)):
            multiplier = raw_multiplier + 1
            score += (multiplier * card)
        return score

    def draw(self):
        return self._deck.popleft()

    def stash(self, card_1, card_2):
        for card in (card_1, card_2):
            self._deck.append(card)

    def copy(self, card_no):
        return Player(self.deck[:card_no])


class Combat:

    log = []

    def __init__(self, player_1, player_2, mode='normal'):
        self.player_1 = player_1
        self.player_2 = player_2
        assert mode in ('normal', 'recursive')
        self.mode = mode
        self._deck_cache = set()
        self._first_round = True

    @staticmethod
    def _process_raw_player(raw_player):
        raw_player_lines = raw_player.splitlines()
        player_name, player_deck = raw_player_lines[:1], raw_player_lines[1:]
        return Player([int(card) for card in player_deck])

    @classmethod
    def from_text(cls, text, mode='normal'):
        raw_player_1, raw_player_2 = text.split('\n\n')
        player_1 = cls._process_raw_player(raw_player_1)
        player_2 = cls._process_raw_player(raw_player_2)
        return cls(player_1, player_2, mode)

    @classmethod
    def from_file(cls, mode='normal'):
        with open('cmd/22/in.txt') as f:
            return cls.from_text(f.read().strip(), mode)

    def _check_previous_decks(self):
        immutable_deck_state = (self.player_1.deck, self.player_2.deck)
        if immutable_deck_state in self._deck_cache:
            self.player_1.win_state = True
            raise StopIteration(f'Player 1 has won the game!')
        self._deck_cache.add(immutable_deck_state)

    def _check_current_decks(self):
        if not self.player_1.deck:
            self.player_2.win_state = True
            raise StopIteration(f'Player 2 has won the game!')
        if not self.player_2.deck:
            self.player_1.win_state = True
            raise StopIteration(f'Player 1 has won the game!')

    def __iter__(self):
        return self

    def __next__(self):
        self._check_previous_decks()
        self._check_current_decks()
        self.play_round()

    def play_game(self):
        for _ in self:
            try:
                next(self)
            except StopIteration:
                return

    def play_round(self):
        if self._first_round:
            Combat.log = []
            self._first_round = False
        Combat.log.append((self.player_1.deck, self.player_2.deck))
        player_1_card = self.player_1.draw()
        player_2_card = self.player_2.draw()
        if (
                self.mode == 'recursive'
                and (self.player_1.cards_remaining >= player_1_card)
                and (self.player_2.cards_remaining >= player_2_card)
        ):
            sub_game_player_1 = self.player_1.copy(player_1_card)
            sub_game_player_2 = self.player_2.copy(player_2_card)
            sub_game = Combat(sub_game_player_1, sub_game_player_2, mode='recursive')
            sub_game._first_round = False
            sub_game.play_game()
            if sub_game.player_1.win_state:
                self.player_1.stash(player_1_card, player_2_card)
            elif sub_game.player_2.win_state:
                self.player_2.stash(player_2_card, player_1_card)
            else:
                raise Exception('One player must win the sub-game.')
        elif player_1_card > player_2_card:
            self.player_1.stash(player_1_card, player_2_card)
        elif player_1_card < player_2_card:
            self.player_2.stash(player_2_card, player_1_card)
        else:
            raise Exception('Two players cannot have the same card.')


def main():
    normal_combat = Combat.from_file(mode='normal')
    normal_combat.play_game()
    print(f'Score in normal mode: {normal_combat.player_1.score + normal_combat.player_2.score}')

    recursive_combat = Combat.from_file(mode='recursive')
    recursive_combat.play_game()
    print(f'Score in recursive mode: {recursive_combat.player_1.score + recursive_combat.player_2.score}')


if __name__ == '__main__':
    print(f'Completed in {timeit.timeit(main, number=1)} seconds')