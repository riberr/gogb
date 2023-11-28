## tutorials:

- https://robertovaccari.com/blog/2020_09_26_gameboy/
- https://rylev.github.io/DMG-01/public/book/cpu/registers.html
- https://github.com/rockytriton/LLD_gbemu/blob/main/part2-3/include/cart.h
- http://www.codeslinger.co.uk/pages/projects/gameboy/lcd.html

## docs:

- https://izik1.github.io/gbops/index.html
- https://gbdev.gg8.se/wiki/articles/Gameboy_Bootstrap_ROM
- https://gbdev.io/pandocs/
- https://gekkio.fi/files/gb-docs/gbctr.pdf
- https://github.com/wheremyfoodat/Gameboy-logs

## resources:

- https://gbdev.gg8.se/files/roms/blargg-gb-tests/    test roms
- https://github.com/wheremyfoodat/Gameboy-logs/tree/master   test rom logs
- https://bgb.bircd.org/  debugger
- https://github.com/raddad772/jsmoo/blob/main/system/gb/gb.js
- https://github.com/Humpheh/goboy/tree/master  main go inspiration
- https://github.com/guigzzz/GoGB/blob/master/backend/cpu.go#L62  go
- https://github.com/theinternetftw/dmgo/blob/master/dmgo.go#L492 go
- https://github.com/duysqubix/gobc go (pixel)
- https://github.com/adtennant/sm83-test-data/tree/master json cpu unit tests
## timing related:

- http://blog.rekawek.eu/2017/02/09/coffee-gb/#cpu-timing
- https://blog.tigris.fr/2021/07/28/writing-an-emulator-timing-is-key/

### repos that do timing by 'stepping' instructions
- https://github.com/rvaccarim/FrozenBoy/blob/dac3dac1d33301019c02a78f9473f80d07999747/FrozenBoyCore/Processor/CPU.cs#L65
- https://github.com/lazy-stripes/goholint/blob/657eaca119e2215dd5bd69b9dd2168ae3a280cd2/cpu/cpu.go#L58
- https://github.com/trekawek/coffee-gb/blob/fb24b380da437890e0447d8b93d11b09d126f38f/src/main/java/eu/rekawek/coffeegb/cpu/Cpu.java#L66

Notes
- GoBoy passes the same gbmicrotests, at least for timer and interrupts
- Mooneye test requires graphics to be implemented or else they just "time out"
- pixel on windows: https://github.com/gopxl/pixel/wiki/Building-Pixel-on-Windows

TODO
- the issue with timing was CB prefix!!
- Try to use Timer(1) with CB fix, cuz that timer passes all the timer microtests