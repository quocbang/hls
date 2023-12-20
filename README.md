# hls

## Cách tính width của chất lượng cần convert

  params: 

    - height_need_convert: number
    - width_of_origin_video: number
    - height_of_origin_video: number
  
  ```

  width = trunc(height_need_convert * width_of_origin_video / height_of_origin_video)

  ```
