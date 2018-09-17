Small script written in go to download images from urls listed in a file in parallel
Command usage is as shown:

File with urls must have a url on each line and nothing more. Must be no extra lines after the final url

Usage of ./imagedownloader:
  -destination string
        The path of where to save the images
  -filename string
        The path of the file with all of the urls
  -timeout int
        Seconds before all unfinished downloads are killed (default 180)