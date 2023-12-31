ffmpeg -i mo_denvau.mp4  \
-map 0:v:0 -map 0:a:0 \
-map 0:v:0 -map 0:a:0 \
-map 0:v:0 -map 0:a:0 \
-map 0:v:0 -map 0:a:0 \
-map 0:v:0 -map 0:a:0 \
-map 0:v:0 -map 0:a:0 \
-c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -c:a aac -ar 48000 \
-filter:v:0 scale="trunc(oh*a/2)*2:240" -maxrate:v:0 800k -bufsize:v:0 400k -b:a:0 84k \
-filter:v:1 scale="trunc(oh*a/2)*2:360" -maxrate:v:1 1600k -bufsize:v:1 800k -b:a:1 96k \
-filter:v:2 scale="trunc(oh*a/2)*2:480" -maxrate:v:2 2100k -bufsize:v:2 1200k -b:a:2 128k \
-filter:v:3 scale="trunc(oh*a/2)*2:720" -maxrate:v:3 3200k -bufsize:v:3 2500k -b:a:3 128k \
-filter:v:4 scale="trunc(oh*a/2)*2:1080" -maxrate:v:4 10000k -bufsize:v:4 5000k -b:a:4 192k \
-filter:v:5 scale="trunc(oh*a/2)*2:1440" -maxrate:v:5 16000k -bufsize:v:5 80000k -b:a:5 192k \
-var_stream_map "v:0,a:0,name:240p v:1,a:1,name:360p v:2,a:2,name:480p v:3,a:3,name:720p v:4,a:4,name:1080p v:5,a:5,name:1440p" \
-hls_time 4 -hls_list_size 0 -master_pl_name master.m3u8 -hls_segment_filename %v/mo_denvau_%06d.ts %v/mo_denvau.m3u8

# check whether video has audio
ffprobe -v error -select_streams a:0 -show_entries stream=codec_type -of default=nw=1:nk=1 mo_denvau.mp4

# check bitrate of video
ffprobe -v error -select_streams v:0 -show_entries stream=bit_rate -of default=nw=1:nk=1 mo_denvau.mp4

# check video width, height
ffprobe -v error -select_streams v:0 -show_entries stream=width,height -of default=nw=1:nk=1 mo_denvau.mp4
