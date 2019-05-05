package quic

//go:generate sh -c "./mockgen_private.sh quic mock_stream_internal_test.go github.com/attenberger/quic-go streamI"
//go:generate sh -c "./mockgen_private.sh quic mock_crypto_stream_test.go github.com/attenberger/quic-go cryptoStream"
//go:generate sh -c "./mockgen_private.sh quic mock_receive_stream_internal_test.go github.com/attenberger/quic-go receiveStreamI"
//go:generate sh -c "./mockgen_private.sh quic mock_send_stream_internal_test.go github.com/attenberger/quic-go sendStreamI"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_sender_test.go github.com/attenberger/quic-go streamSender"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_getter_test.go github.com/attenberger/quic-go streamGetter"
//go:generate sh -c "./mockgen_private.sh quic mock_crypto_data_handler_test.go github.com/attenberger/quic-go cryptoDataHandler"
//go:generate sh -c "./mockgen_private.sh quic mock_frame_source_test.go github.com/attenberger/quic-go frameSource"
//go:generate sh -c "./mockgen_private.sh quic mock_ack_frame_source_test.go github.com/attenberger/quic-go ackFrameSource"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_manager_test.go github.com/attenberger/quic-go streamManager"
//go:generate sh -c "./mockgen_private.sh quic mock_sealing_manager_test.go github.com/attenberger/quic-go sealingManager"
//go:generate sh -c "./mockgen_private.sh quic mock_unpacker_test.go github.com/attenberger/quic-go unpacker"
//go:generate sh -c "./mockgen_private.sh quic mock_packer_test.go github.com/attenberger/quic-go packer"
//go:generate sh -c "./mockgen_private.sh quic mock_session_runner_test.go github.com/attenberger/quic-go sessionRunner"
//go:generate sh -c "./mockgen_private.sh quic mock_quic_session_test.go github.com/attenberger/quic-go quicSession"
//go:generate sh -c "./mockgen_private.sh quic mock_packet_handler_test.go github.com/attenberger/quic-go packetHandler"
//go:generate sh -c "./mockgen_private.sh quic mock_unknown_packet_handler_test.go github.com/attenberger/quic-go unknownPacketHandler"
//go:generate sh -c "./mockgen_private.sh quic mock_packet_handler_manager_test.go github.com/attenberger/quic-go packetHandlerManager"
//go:generate sh -c "./mockgen_private.sh quic mock_multiplexer_test.go github.com/attenberger/quic-go multiplexer"
