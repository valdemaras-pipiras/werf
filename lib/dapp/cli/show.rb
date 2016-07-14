require 'mixlib/cli'

module Dapp
  class CLI
    # CLI show subcommand
    class Show < Base
      banner <<BANNER.freeze
Version: #{Dapp::VERSION}

Usage:
  dapp show [options] [PATTERN ...]

    PATTERN                     Applications to process [default: *].

Options:
BANNER
    end
  end
end
