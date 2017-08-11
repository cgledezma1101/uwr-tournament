class TournamentTeam < ActiveRecord::Base
	attr_accessor :password_confirmation

	belongs_to :tournament
	belongs_to :team

	validate :club_invited
	validate :passwords_match

	validate :team, uniqueness: { scope: :tournament }
	validate :password

	before_save :scramble_password

	def valid_password?(test)
		scrambled_test = Digest::MD5.hexdigest(test)
		return scrambled_test == self.password
	end

	private

	# Creates a hash of the password to prevent saving verbatim string to the DB
	def scramble_password
		if self.password.changed?
			self.password = Digest::MD5.hexdigest(self.password)
		end
	end

	# Makes sure that the club the team belongs to has been invited to this tournament
	def club_invited
		self.errors.add(:team, I18n.t('tournament.team_uninvited')) unless self.tournament.clubs.where{ id == my{ team.club.id } }.any?
	end

	# When the password has changed, validates that the password confirmation has the same value
	def passwords_match
		if self.password.changed? && !(self.password == self.password_confirmation)
			self.errors.add(:team, I18n.t('team.passwords_dont_match'))
		end
	end
end
