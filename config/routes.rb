UwrTournament::Application.routes.draw do
  root 'games#new'

  devise_for :users

  get '/statistics', to: 'public#statistics'

  resources :clubs, only: [:show, :new, :create, :edit, :update]

  resources :games, only: [:create, :new, :show, :index]

  resources :scores, only: [:create]

  resources :teams, only: [:new, :create, :show] do
    get 'players', on: :member
  end

  resources :invitations, only: [:new, :create] do
    get 'accept', on: :member
    get 'decline', on: :member
  end

  resources :club_admins, only: [:new, :create, :destroy] do
    get 'confirm_destroy', on: :member
  end
end
