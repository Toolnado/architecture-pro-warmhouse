package com.warmhouse.device.model;

import jakarta.persistence.*;
import java.util.UUID;
import lombok.Data;

@Entity
@Table(name = "devices")
@Data
public class Device {

  @Id
  @GeneratedValue(strategy = GenerationType.UUID)
  private UUID id;

  @Column(nullable = false)
  private String name;

  @Column(nullable = false)
  private String type;

  @Column(name = "serial_number", nullable = false)
  private String serialNumber;

  @Column(nullable = false)
  private String status = "active";
}
