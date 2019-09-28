class Invitation < ApplicationRecord
	belongs_to :club
	belongs_to :user

	validates :club, presence: true
	validates :user, presence: true

	validate :one_invitation_per_club
	validate :cant_invite_members
	validate :cant_invite_requested

	# Determines the name of the club that generated the invitation
	#
	# @return [String] Name of the associated club
	def club_name
		self.club.name
	end

	private

	# Validates the uniqueness of the user in the scope of the club
	def one_invitation_per_club
		if !self.user.nil? && !self.club.nil?
			duplicate = Invitation.find_by(club: self.club, user: self.user)

			if !duplicate.nil? && duplicate.id != self.id
				self.errors.add(:user, I18n.t('invitation.only_one_per_club'))
			end
		end
	end

	# Validates that an invitation is not sent to a person that's already part of the club
	def cant_invite_members
		if !self.user.nil? && !self.club.nil?
			is_member = self.club.members.any?{ |member| member.id == self.user.id }

			if is_member
				self.errors.add(:user, I18n.t('invitation.cant_invite_members'))
			end
		end
	end

	# Validates that an invitation is not sent to someone who has a pending membership request
	def cant_invite_requested
		if ClubJoinRequest.where(club_id: self.club_id, user_id: self.user_id).any?
			self.errors.add(:user, I18n.t('invitation.cant_invite_requested'))
		end
	end
end
