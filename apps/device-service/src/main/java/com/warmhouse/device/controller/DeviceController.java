package com.warmhouse.device.controller;

import com.warmhouse.device.model.Device;
import com.warmhouse.device.service.DeviceService;
import java.util.List;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/devices")
public class DeviceController {

  private final DeviceService deviceService;

  public DeviceController(DeviceService deviceService) {
    this.deviceService = deviceService;
  }

  @PostMapping
  public ResponseEntity<Device> registerDevice(@RequestBody Device device) {
    return ResponseEntity.status(HttpStatus.CREATED).body(
      deviceService.registerDevice(device)
    );
  }

  @GetMapping
  public ResponseEntity<List<Device>> getAllDevices() {
    return ResponseEntity.ok(deviceService.getAllDevices());
  }

  @GetMapping("/health")
  public ResponseEntity<String> health() {
    return ResponseEntity.ok("Device Service is running");
  }
}
