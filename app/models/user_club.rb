class UserClub < ApplicationRecord
	belongs_to :user
	belongs_to :club

	validates :user, presence: true, uniqueness: { scope: :club }
	validates :club, presence: true

	# The full name of the user included in this relationship
	#
	# @return [String] Full name of the user associated
	def user_name
		self.user.name
	end

	# The name of the club associated
	#
	# @return [String] Name of the associated club
	def club_name
		self.club.name
	end
end
