module main
	require "log"
	require "net/http"

	def main()
		fs = http.FileServer(http.Dir("static"))
		http.Handle("/", fs)

		puts "listening"
		http.
	end
end
