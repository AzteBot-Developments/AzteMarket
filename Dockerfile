FROM rust:latest

COPY ./aztemarket ./

RUN cargo build --release

CMD ["./target/release/aztemarket"]