Comparison of serving image files

This repository is a performance comparison of `http.Dir()` and `//go:embed` for serving image files.

## Target image files

```
% ls -1 images | wc
     744

% du -h images
 52M    images

# This means that the average size of images is about 70K.
```

## Result

When using `http.Dir`

```

% go build -o client-bin ./client
% time ./client-bin
./client-bin  17.03s user 20.65s system 81% cpu 46.225 total
```

When using `//go:embed`

```

% time ./client-bin
./client-bin  10.79s user 11.32s system 86% cpu 25.706 total
```
# compare-serve-images
